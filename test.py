import os
from dotenv import load_dotenv
from requests_html import HTMLSession
from bs4 import BeautifulSoup

load_dotenv()

def scrape_dynamic_content(url):
    session = HTMLSession()
    r = session.get(url)
    r.html.render()

    soup = BeautifulSoup(r.html.html, "html.parser")

    data = soup.find_all("article", class_="table_content__53NAZ")
    print(data)
    for d in data:
        pass

    return None

# url = os.getenv("BEEFY_URL")
url = "https://debank.com/profile/0x1c45e086ed143aef83c1209521a2ff5369f39abc"
dynamic_content = scrape_dynamic_content(url)
