import json
from typing import Optional

import uvicorn
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def read_root():
    return {"hello": "World"}

@app.get("/recipes")
def get_recipes():
    json_data = open(file="recipes.json")
    data = json.load(json_data)
    return data

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8001)