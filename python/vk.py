import argparse
import json
import requests


def main():
    parser = argparse.ArgumentParser(description='VK API getter')
    parser.add_argument('-t', '--token', action='store', dest="token", type=str, help='VK API Token.')
    token = parser.parse_args().token

    get_wall = get_vk(token, "wall.get", {
        "domain": "mudakoff",
        "count": 1,
    })

    owner_id = get_wall["response"]["items"][0]["owner_id"]

    get_wall_comments = get_vk(token, "wall.getComments", {
        "owner_id": owner_id,
        "post_id": 36857907,
        "count": 100,
        "thread_items_count": 10,
    })

    result = [
        {
            "author": item.get("from_id", 0),
            "text": item.get("text", ""),
            "replies": [
                {
                    "author": jtem.get("from_id", 0),
                    "text": jtem.get("text", "")
                }
                for jtem in item["thread"].get("items", [])
            ]
        }
        for item in get_wall_comments["response"]["items"]
    ]

    with open('output/vk_py.json', 'w') as target:
        target.write(json.dumps(result, ensure_ascii=False))


def get_vk(token, method, values):
    values["access_token"] = token
    values["v"] = "5.91"
    response = requests.get('https://api.vk.com/method/' + method, params=values)
    return json.loads(response.text)


main()
