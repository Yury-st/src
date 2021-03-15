import csv, requests, time

def take_posts():

    token = 'd5b58fd5d5b58fd5d5b58fd5e7d5c34a61dd5b5d5b58fd5b5f08044afe0742b6c884c75'
    version = '5.130'
    domain = 'adkazy'
    count = 100
    offset = 0
    all_posts = []

    while offset < 20000:
        response = requests.get('https://api.vk.com/method/wall.get',
                            params={
                                'access_token': token,
                                'v': version,
                                'domain': domain,
                                'count': count,
                                'offset': offset
                            }
    )
        data = response.json()['response']['items']
        offset += 100
        all_posts.extend(data)
    return all_posts


def file_writer(data):
    with open('vk.csv', 'w') as file:
        a_pen = csv.writer(file)
        a_pen.writerow(('likes', 'body'))
        for post in data:
            # try:
            #     if post['attachments'][0]['type']:
            #         img_url = post['attachments'][0]['photo']['sizes'][-1]['url']
            #     else:
            #         img_url = 'pass'
            # except:
            #     pass
            a_pen.writerow((post['likes']['count'], post['text']))

all_posts = take_posts()
file_writer(all_posts)