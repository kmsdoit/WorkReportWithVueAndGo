from fastapi import FastAPI

from database import session
from database.session import engine

session.Base.metadata.create_all(bind=engine)
app = FastAPI()


@app.get("/ping")
def pong():
    return 'pong'
