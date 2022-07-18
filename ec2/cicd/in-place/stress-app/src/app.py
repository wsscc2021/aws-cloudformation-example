# Standard Library
import time
import itertools
# Application modules
from flask import Flask
from healthcheck import HealthCheck

# Create Flask App
app = Flask(__name__)

# Healthcheck API
health = HealthCheck()
app.add_url_rule("/healthcheck", "healthcheck", view_func=lambda: health.run())

# nothing, it just say hello.
@app.route("/greet")
def get_greet():
    return "Hello World!", 200

# cpu
@app.route("/cpu")
def get_cpu():
    for i in range(1,2500):
        i ** i
    return "CPU load...", 200

# memory
@app.route("/memory")
def get_memory():
    dummy = list(itertools.repeat(['A'],10000000))
    time.sleep(0.2)
    return "Memory load...", 200

# limit connection
count=0
@app.route("/connection")
def get_connection():
    global count
    count += 1
    if count > 5: time.sleep(3)
    else: time.sleep(0.1)
    count -= 1
    return "Connection lazy...", 200

if __name__ == "__main__":
    RUN_OPTIONS = {"host": "0.0.0.0", "port": 5000, "threaded": True}
    app.run(**RUN_OPTIONS)