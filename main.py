import threading

api_key = ""
api_secret = ""

def printit():
    threading.timer(1, 0, printit).start()
