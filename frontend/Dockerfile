# ベースはNode.js
FROM node:22

WORKDIR /usr/src/app

RUN apt update && apt install -y bash

# COPY ./package*.json .
# RUN npm install

# COPY ./project . # nodemodulesをコピーしようとして問題が発生した．

EXPOSE 3000

# CMD ["npm", "start"]
CMD ["tail", "-f", "/dev/null"]