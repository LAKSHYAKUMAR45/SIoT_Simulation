import os
import torch
from flask import Flask, request, jsonify
from transformers import DistilBertForSequenceClassification, DistilBertTokenizer

app = Flask(__name__)

# Load the trained model
model_path = "model_converted_dataset_11.csv.pth"
model = DistilBertForSequenceClassification.from_pretrained('distilbert-base-uncased', num_labels=2)
model.load_state_dict(torch.load(model_path))
model.eval()

# Load the tokenizer
tokenizer = DistilBertTokenizer.from_pretrained('distilbert-base-uncased')

# Define a function to predict
def predict(text):
    inputs = tokenizer(text, return_tensors="pt", padding=True, truncation=True)
    outputs = model(**inputs)
    logits = outputs.logits
    predicted_class = torch.argmax(logits, dim=1).item()
    return predicted_class

# Define a route for prediction
@app.route('/predict', methods=['POST'])
def get_prediction():
    data = request.get_json(force=True)
    text = data['text']
    prediction = predict(text)
    return jsonify({'prediction': prediction})

if __name__ == '__main__':
    app.run(debug=True)
