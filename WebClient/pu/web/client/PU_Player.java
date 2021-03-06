package pu.web.client;

public class PU_Player extends PU_Creature
{
	public static final int NUM_BODYPARTS = 6;
	public static final int NUM_POKEMON = 6;
	
	public static final int BODY_BASE  = 0;
	public static final int BODY_UPPER = 1;
	public static final int BODY_NECK  = 2;
	public static final int BODY_HEAD  = 3;
	public static final int BODY_FEET  = 4;
	public static final int BODY_LOWER = 5;
	
	private long mMoney = 0;
	
	private PU_BodyPart[] mBodyParts = new PU_BodyPart[NUM_BODYPARTS];
	
	private PU_Pokemon[] mPokemon = new PU_Pokemon[6];
	
	public PU_Player(long id)
	{
		mId = id;
		
		for(int i = 0; i < NUM_BODYPARTS; i++)
		{
			mBodyParts[i] = new PU_BodyPart(1);
		}
	}
	
	public void setMoney(long money)
	{
		mMoney = money;
	}
	
	public long getMoney()
	{
		return mMoney;
	}
	
	public void turn(int direction, boolean send)
	{
		if(direction != mDirection)
		{
			mDirection = direction;
			
			if(send)
			{
				PUWeb.connection().getProtocol().sendTurn(direction);
			}
		}
	}
	
	public void draw(int x, int y)
	{
		for(int part = BODY_BASE; part < BODY_LOWER; part++)
		{
			PU_Image image = PUWeb.resources().getCreatureImage(part, mBodyParts[part].id, mDirection, mFrame);
			if(image != null)
			{
				PUWeb.engine().addToSpriteBatch(image, x, y);
			}
		}
	}
	
	public void setBodyPart(int part, int id)
	{
		mBodyParts[part].id = id;
	}
	
	public PU_BodyPart getBodyPart(int part)
	{
		return mBodyParts[part];
	}
	
	public void walk(int direction)
	{
		walk(direction, false);
	}
	
	public void walk(int direction, boolean force)
	{
		if(!mWalking || force)
		{
			if(preWalk(direction, false))
			{
				PUWeb.connection().getProtocol().sendWalk(direction, true);
			}
			else
			{
				cancelWalk();
			}
		}
	}
	
	public boolean preWalk(int direction, boolean continuing)
	{
		PU_Tile toTile = null;
		switch(direction)
		{
		case PU_Creature.DIR_NORTH:
			toTile = PUWeb.map().getTile(mX, mY-1);
			break;
			
		case PU_Creature.DIR_EAST:
			toTile = PUWeb.map().getTile(mX+1, mY);
			break;
			
		case PU_Creature.DIR_SOUTH:
			toTile = PUWeb.map().getTile(mX, mY+1);
			break;
			
		case PU_Creature.DIR_WEST:
			toTile = PUWeb.map().getTile(mX-1, mY);
			break;
		}
		
		if(canWalkTo(direction, toTile))
		{
			mPreWalkX = toTile.getX();
			mPreWalkY = toTile.getY();
			
			if(continuing && direction == mDirection)
			{
				mWalkProgress = mWalkProgress - 1.0f;
				mOffset = (int)Math.round(mWalkProgress * (float)PU_Tile.TILE_WIDTH);
			}
			else
			{
				turn(direction, false);
				
				if(!mAnimationRunning)
				{
					startAnimation();
				}
				
				mWalkProgress = 0.0f;
				mOffset = 0;
				
				mWalking = true;
			}
			return true;
		}
		return false;
	}
	
	public boolean canWalkTo(int direction, PU_Tile tile)
	{
		if(tile == null)
		{
			return false;
		}
		
		int tileMovement = tile.getMovement();
		if(tileMovement != PU_Tile.TILE_WALK) 
		{
			if((tileMovement == PU_Tile.TILE_BLOCKING) ||
				(tileMovement == PU_Tile.TILE_SURF) ||
				(tileMovement == PU_Tile.TILE_BLOCKTOP && direction == PU_Creature.DIR_SOUTH) ||
				(tileMovement == PU_Tile.TILE_BLOCKBOTTOM && direction == PU_Creature.DIR_NORTH) ||
				(tileMovement == PU_Tile.TILE_BLOCKLEFT && direction == PU_Creature.DIR_EAST) ||
				(tileMovement == PU_Tile.TILE_BLOCKRIGHT && direction == PU_Creature.DIR_WEST) ||
				(tileMovement == PU_Tile.TILE_BLOCKCORNER_TL && (direction == PU_Creature.DIR_EAST || direction == PU_Creature.DIR_SOUTH)) ||
				(tileMovement == PU_Tile.TILE_BLOCKCORNER_TR && (direction == PU_Creature.DIR_WEST || direction == PU_Creature.DIR_SOUTH)) ||
				(tileMovement == PU_Tile.TILE_BLOCKCORNER_BL && (direction == PU_Creature.DIR_EAST || direction == PU_Creature.DIR_NORTH)) ||
				(tileMovement == PU_Tile.TILE_BLOCKCORNER_BR && (direction == PU_Creature.DIR_WEST || direction == PU_Creature.DIR_NORTH)))
			{
				return false;
			}
		}
		return true;
	}
	
	public void cancelWalk()
	{
		mWalking = false;
		mWalkProgress = 0.0f;
		mOffset = 0;
		stopAnimation();
	}
	
	public void receiveWalk(PU_Tile fromTile, PU_Tile toTile)
	{
		if(toTile == null || fromTile == null)
		{
			cancelWalk();
			return;
		}
		
		if(mX != fromTile.getX() || mY != fromTile.getY())
		{
			mX = fromTile.getX();
			mY = fromTile.getY();
			
			mPreWalkX = toTile.getX();
			mPreWalkY = toTile.getY();
		}
		
		if(this != PUWeb.game().getSelf())
		{
			mPreWalkX = toTile.getX();
			mPreWalkY = toTile.getY();
			
			if(mPreWalkY > mY)
			{
				turn(DIR_SOUTH, false);
			}
			else if(mPreWalkY < mY)
			{
				turn(DIR_NORTH, false);
			}
			else if(mPreWalkX > mX)
			{
				turn(DIR_EAST, false);
			}
			else if(mPreWalkX < mX)
			{
				turn(DIR_WEST, false);
			}
			mWalkProgress = 0.0f;
			mOffset = 0;
			startAnimation();
			mWalking = true;
		}
		else
		{
			mX = mPreWalkX;
			mY = mPreWalkY;
			
			cancelWalk();
		}
	}
	
	public void updateWalk()
	{
		if(mWalking)
		{
			mWalkProgress += (1000.0f / (float)mSpeed) * ((float)PUWeb.getFrameTime() / 1000.0f);
			if(mWalkProgress >= 1.0f)
			{
				endWalk();
			}
			else
			{
				mOffset = (int)Math.round(mWalkProgress * (float)PU_Tile.TILE_WIDTH);
			}
			updateAnimation();
		}
	}
	
	public void endWalk()
	{
		if(PUWeb.game().getSelf() == this)
		{			
			mX = mPreWalkX;
			mY = mPreWalkY;
			if(!continueWalk())
			{	
				stopAnimation();
				mWalking = false;
			}
		}
		else
		{
			mOffset = PU_Tile.TILE_WIDTH;
			
			mX = mPreWalkX;
			mY = mPreWalkY;
			
			mWalking = false;
			stopAnimation();
		}
	}
	
	public boolean continueWalk()
	{
		if(PUWeb.game().getState() != PU_Game.GAMESTATE_WORLD)
		{
			return false;
		}
		
		if(this == PUWeb.game().getSelf())
		{
			if(PUWeb.events().isKeyDown(PUWeb.game().getLastDirKey()))
			{
				boolean continuing = false;
				switch(PUWeb.game().getLastDirKey())
				{
				case PU_Events.KEY_UP:
					continuing = preWalk(DIR_NORTH, true);
					break;
					
				case PU_Events.KEY_DOWN:
					continuing = preWalk(DIR_SOUTH, true);
					break;
					
				case PU_Events.KEY_LEFT:
					continuing = preWalk(DIR_WEST, true);
					break;
					
				case PU_Events.KEY_RIGHT:
					continuing = preWalk(DIR_EAST, true);
					break;
				}
				if(continuing)
				{
					PUWeb.connection().getProtocol().sendWalk(mDirection, true);
				}
				return continuing;
			}
		}
		return false;
	}
	
	public void setPokemon(int slot, PU_Pokemon pokemon)
	{
		this.mPokemon[slot] = pokemon;
	}
	
	public PU_Pokemon getPokemon(int slot)
	{
		return this.mPokemon[slot];
	}
	
	public int getPokemonCount()
	{
		int count = 0;
		for(int i = 0; i < 6; i++)
		{
			if(this.mPokemon[i] != null)
				count++;
		}
		return count;
	}
}
