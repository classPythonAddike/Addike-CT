from flask import Flask, jsonify
from flask_cors import CORS, cross_origin

app = Flask(__name__)
CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

starter_data = {
    "name": "Simon Says",
    "description": "Only what Simon said counts! On each test, you will be given a  string of any length. This string can be anything. Your job is to output what Simon said!",
    "files": [{
        "name": "123456789.py",
        "progress": 100,
        "length": 45
    },
    {
        "name": "123456789.py",
        "progress": 75,
        "length": 56
    },
    {
        "name": "123456789.py",
        "progress": 50,
        "length": 78
    },
    {
        "name": "123456789.py",
        "progress": 25,
        "length": 90
    },
    {
        "name": "123456789.py",
        "progress": 1,
        "length": 183
    }]
}
progresses = {"values": [99,74,49,24,2],
                "times": [29.34,29.34,29.34,29.34,29.34]}

@app.route('/startup')
@cross_origin()
def startup():
    return jsonify(starter_data)

@app.route('/progress')
@cross_origin()
def progress():
    for progress in progresses["values"]:
        index = progresses["values"].index(progress)
        if progress > 100:
            progresses["values"][index] = 0
        progresses["values"][index] += 1
    return jsonify(progresses)

if __name__ == '__main__':
    app.run()