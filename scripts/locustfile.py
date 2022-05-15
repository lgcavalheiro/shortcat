
from locust import task
from locust.contrib.fasthttp import FastHttpUser
from random import choice

TARGETS = ["LH11aA", "wG7VhO", "NChHns", "x8G9Ph", "JMCyPd", "ia5PXZ"]


class QuickstartUser(FastHttpUser):
    @task
    def get_shortcat(self):
        self.client.get(f"/go/{choice(TARGETS)}")
