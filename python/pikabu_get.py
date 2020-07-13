import urllib.request

from bs4 import BeautifulSoup

response = urllib.request.urlopen('https://pikabu.ru/story/sluzhba_bezopasnosti_banka_i_teamviewer_7579039')
soup = BeautifulSoup(response.read(), 'html.parser')

with open('pikabu.html', 'w') as target:
    target.write(soup.prettify())
