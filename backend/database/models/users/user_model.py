from sqlalchemy import Column, String, Integer, ForeignKey
from database.session import Base


class User(Base):
    __tablename__ = "users"

    no = Column(Integer, primary_key=True, index=True)
    email = Column(String, nullable=False)
    password = Column(String)
    name = Column(String)
    dept = Column(String)
