import tkinter as tk
import requests
import json
from PIL import Image, ImageTk
from io import BytesIO
import quickLLM as summarizer

#To get a synopsis of the manga, I am going to add an LLM to get lebron james to summarize this in his own
def getSynopsis(msg, image, search_bar, root):
    dataName = {"mangaName": search_bar.get()}

    response = requests.post("http://localhost:3333/skibidiRizzlerSigmaMale", json=dataName) #get response
    data = response.json()


    updateImage(data["imageurl"], image)
    summarizer.summarizeMangaSynopsis(data["response"], msg, root)
    search_bar.set("")

# #To get score, I am going to get an LLM to summarize this, so it would seem cooler!
def getScore(msg, image, search_bar):
    dataName = {"mangaName": search_bar.get()}
    response = requests.post("http://localhost:3333/getScore", json=dataName)
    data = response.json()

    msg.config(text=data["response"])
    updateImage(data["imageurl"], image)
    search_bar.set("")

#To get a recommendation from the given thing
def getRecommendation(msg, image, search_bar):
    dataName = {"mangaName": search_bar.get()}
    response = requests.post("http://localhost:3333/getRec", json=dataName)
    data = response.json()

    response = "the recommended manga:\n" + data["response"]
    msg.config(text=response)
    updateImage(data["imageurl"], image)
    search_bar.set("")

#Updates the image per search
def updateImage(url, imageFrame):
    if len(url) > 0: 
        resp = requests.get(url)
        image = Image.open(BytesIO(resp.content))
        photo = ImageTk.PhotoImage(image)
        imageFrame.config(image=photo)  #shows the image
        imageFrame.image = photo   #keeps the image alive in the manga_photo attribute, to avoid garbage collector
    else: 
        print("no images i guess")

import ollama
