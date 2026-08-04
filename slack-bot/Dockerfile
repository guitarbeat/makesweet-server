FROM python:3.11-slim

COPY requirements.txt /app/requirements.txt
RUN pip install --no-cache-dir -r /app/requirements.txt

COPY bot.py /app/bot.py

WORKDIR /app
EXPOSE 3000
CMD ["python", "bot.py"]
