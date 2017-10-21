import sys
sys.path.append('../../') # NoQA
from n0core.lib.proto import CreateVolumeRequest, DeleteVolumeRequest
from n0core.lib.n0mq import N0MQ


def main():
    client = N0MQ('pulsar://localhost:6650')
    producer = client.create_producer('persistent://sample/standalone/volumer/114514')

    req = CreateVolumeRequest(id='1', host='hoge', size_mb=1)
    producer.send(req)

    req = DeleteVolumeRequest(id='1', host='hoge')
    producer.send(req)

    client.close()


if __name__ == '__main__':
    main()