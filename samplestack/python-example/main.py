from typing import List, Optional

from fastapi import FastAPI
import uvicorn
from pydantic import BaseModel

app = FastAPI()


class Item(BaseModel):
    name: str
    description: Optional[str] = None
    price: float
    tax: Optional[float] = None
    tags: List[str] = []


@app.get("/api/v1/item", response_model=Item)
async def get_item():
    item = Item(name="pong", price=5.0, tags=["some", "endpoint"], tax=21.0, description="one")
    return item


if __name__ == '__main__':
    uvicorn.run(app, port=5000, host='0.0.0.0')
