import json
import os
from flask import Flask, url_for

app = Flask(__name__)


@app.route('/recipes')
def get_recipes():  # put application's code here
    with open('recipes.json', 'r') as f:
        data = f.read()
    return data

if __name__ == '__main__':
    app.run()
