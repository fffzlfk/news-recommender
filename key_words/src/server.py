from textrank4zh import TextRank4Keyword, TextRank4Sentence
from concurrent import futures
import time
import grpc
import keywords_pb2
import keywords_pb2_grpc
import json
import codecs


# 实现 proto 文件中定义的 GreeterServicer
class Greeter(keywords_pb2_grpc.GreeterServicer):
    # 实现 proto 文件中定义的 rpc 调用
    def GetKeywords(self, request, context):
        return keywords_pb2.GetKeywordsResp(keywords=get_keywords(request.title))


def get_keywords(text):
    tr4w = TextRank4Keyword()

    # py2中text必须是utf8编码的str或者unicode对象，py3中必须是utf8编码的bytes或者str对象
    tr4w.analyze(text=text, lower=True, window=2)

    res = []
    print('关键词：')
    for item in tr4w.get_keywords(5, word_min_len=2):
        print(item.word, item.weight)
        res.append(keywords_pb2.Item(word=item.word, weight=item.weight))

    return res


def serve():
    # 启动 rpc 服务
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    keywords_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    print('INFO: server started')

    try:
        while True:
            time.sleep(60*60*24)  # one day in seconds
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
