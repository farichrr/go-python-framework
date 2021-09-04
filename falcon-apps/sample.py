import json
import falcon.asgi


f = open("recipes.json")

data = json.load(f)

class QuoteResource:
    async def on_get(self, req, resp):
        # resp.status = falcon.HTTP_200
        # load_json = open("recipes.json")
        # data = json.load(load_json)
        # return data

        resp.media = data
        resp.status = 200


app = falcon.asgi.App()

app.add_route("/", QuoteResource())