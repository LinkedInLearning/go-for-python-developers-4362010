from dataclasses import dataclass
from enum import Enum


class Role(Enum):
    VIEWER = 'viewer'
    DEVELOPER = 'developer'
    ADMIN = 'admin'


@dataclass
class User:
    login: str
    role: Role


def promote(user, role):
    user.role = role


if __name__ == '__main__':
    user = User('elliot', Role.VIEWER)
    promote(user, Role.ADMIN)
    print(user)
