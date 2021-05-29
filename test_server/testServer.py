from flask import Flask, jsonify
from flask_cors import CORS, cross_origin

app = Flask(__name__)
CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

starter_data = {
    "Name": "Simon Says",
    "Description": "Only what Simon said counts! On each test, you will be given a  string of any length. This string can be anything. Your job is to output what Simon said!",
    "Files": [{
        "Name": "123456789.py",
        "Progress": 100,
        "Length": 45
    },
    {
        "Name": "123456789.py",
        "Progress": 75,
        "Length": 56
    },
    {
        "Name": "123456789.py",
        "Progress": 50,
        "Length": 78
    },
    {
        "Name": "123456789.py",
        "Progress": 25,
        "Length": 90
    },
    {
        "Name": "123456789.py",
        "Progress": 1,
        "Length": 183
    }]
}
progresses = {"Values": [99,74,49,24,2],
                "Times": [29.34,29.34,29.34,29.34,29.34]}

@app.route('/startup')
@cross_origin()
def startup():
    return jsonify(starter_data)

@app.route('/progress')
@cross_origin()
def progress():
    for progress in progresses["Values"]:
        index = progresses["Values"].index(progress)
        if progress > 100:
            progresses["Values"][index] = 0
        progresses["Values"][index] += 1
    return jsonify(progresses)

if __name__ == '__main__':
    app.run(port=8080)