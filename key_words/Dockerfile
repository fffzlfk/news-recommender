FROM python:3.9.7

RUN mkdir key_words

ADD . /key_words/

WORKDIR /key_words

RUN pip3 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

RUN pip3 install -r /key_words/requirements.txt

EXPOSE 50052

CMD [ "python3", "/key_words/src/server.py" ]