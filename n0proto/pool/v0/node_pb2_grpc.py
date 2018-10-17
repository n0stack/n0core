# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from pool.v0 import node_pb2 as pool_dot_v0_dot_node__pb2


class NodeServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ListNodes = channel.unary_unary(
        '/n0stack.pool.NodeService/ListNodes',
        request_serializer=pool_dot_v0_dot_node__pb2.ListNodesRequest.SerializeToString,
        response_deserializer=pool_dot_v0_dot_node__pb2.ListNodesResponse.FromString,
        )
    self.GetNode = channel.unary_unary(
        '/n0stack.pool.NodeService/GetNode',
        request_serializer=pool_dot_v0_dot_node__pb2.GetNodeRequest.SerializeToString,
        response_deserializer=pool_dot_v0_dot_node__pb2.Node.FromString,
        )
    self.ApplyNode = channel.unary_unary(
        '/n0stack.pool.NodeService/ApplyNode',
        request_serializer=pool_dot_v0_dot_node__pb2.ApplyNodeRequest.SerializeToString,
        response_deserializer=pool_dot_v0_dot_node__pb2.Node.FromString,
        )
    self.DeleteNode = channel.unary_unary(
        '/n0stack.pool.NodeService/DeleteNode',
        request_serializer=pool_dot_v0_dot_node__pb2.DeleteNodeRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.ScheduleCompute = channel.unary_unary(
        '/n0stack.pool.NodeService/ScheduleCompute',
        request_serializer=pool_dot_v0_dot_node__pb2.ScheduleComputeRequest.SerializeToString,
        response_deserializer=pool_dot_v0_dot_node__pb2.Node.FromString,
        )
    self.ReserveCompute = channel.unary_unary(
        '/n0stack.pool.NodeService/ReserveCompute',
        request_serializer=pool_dot_v0_dot_node__pb2.ReserveComputeRequest.SerializeToString,
        response_deserializer=pool_dot_v0_dot_node__pb2.Node.FromString,
        )
    self.ReleaseCompute = channel.unary_unary(
        '/n0stack.pool.NodeService/ReleaseCompute',
        request_serializer=pool_dot_v0_dot_node__pb2.ReleaseComputeRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.ScheduleStorage = channel.unary_unary(
        '/n0stack.pool.NodeService/ScheduleStorage',
        request_serializer=pool_dot_v0_dot_node__pb2.ScheduleStorageRequest.SerializeToString,
        response_deserializer=pool_dot_v0_dot_node__pb2.Node.FromString,
        )
    self.ReserveStorage = channel.unary_unary(
        '/n0stack.pool.NodeService/ReserveStorage',
        request_serializer=pool_dot_v0_dot_node__pb2.ReserveStorageRequest.SerializeToString,
        response_deserializer=pool_dot_v0_dot_node__pb2.Node.FromString,
        )
    self.ReleaseStorage = channel.unary_unary(
        '/n0stack.pool.NodeService/ReleaseStorage',
        request_serializer=pool_dot_v0_dot_node__pb2.ReleaseStorageRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )


class NodeServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def ListNodes(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetNode(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ApplyNode(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteNode(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ScheduleCompute(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ReserveCompute(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ReleaseCompute(self, request, context):
    """rpc ResizeCompute(ResizeComputeRequest) returns (ReserveComputeResponse) {}
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ScheduleStorage(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ReserveStorage(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ReleaseStorage(self, request, context):
    """rpc ResizeStorae() returns () {}
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_NodeServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ListNodes': grpc.unary_unary_rpc_method_handler(
          servicer.ListNodes,
          request_deserializer=pool_dot_v0_dot_node__pb2.ListNodesRequest.FromString,
          response_serializer=pool_dot_v0_dot_node__pb2.ListNodesResponse.SerializeToString,
      ),
      'GetNode': grpc.unary_unary_rpc_method_handler(
          servicer.GetNode,
          request_deserializer=pool_dot_v0_dot_node__pb2.GetNodeRequest.FromString,
          response_serializer=pool_dot_v0_dot_node__pb2.Node.SerializeToString,
      ),
      'ApplyNode': grpc.unary_unary_rpc_method_handler(
          servicer.ApplyNode,
          request_deserializer=pool_dot_v0_dot_node__pb2.ApplyNodeRequest.FromString,
          response_serializer=pool_dot_v0_dot_node__pb2.Node.SerializeToString,
      ),
      'DeleteNode': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteNode,
          request_deserializer=pool_dot_v0_dot_node__pb2.DeleteNodeRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'ScheduleCompute': grpc.unary_unary_rpc_method_handler(
          servicer.ScheduleCompute,
          request_deserializer=pool_dot_v0_dot_node__pb2.ScheduleComputeRequest.FromString,
          response_serializer=pool_dot_v0_dot_node__pb2.Node.SerializeToString,
      ),
      'ReserveCompute': grpc.unary_unary_rpc_method_handler(
          servicer.ReserveCompute,
          request_deserializer=pool_dot_v0_dot_node__pb2.ReserveComputeRequest.FromString,
          response_serializer=pool_dot_v0_dot_node__pb2.Node.SerializeToString,
      ),
      'ReleaseCompute': grpc.unary_unary_rpc_method_handler(
          servicer.ReleaseCompute,
          request_deserializer=pool_dot_v0_dot_node__pb2.ReleaseComputeRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'ScheduleStorage': grpc.unary_unary_rpc_method_handler(
          servicer.ScheduleStorage,
          request_deserializer=pool_dot_v0_dot_node__pb2.ScheduleStorageRequest.FromString,
          response_serializer=pool_dot_v0_dot_node__pb2.Node.SerializeToString,
      ),
      'ReserveStorage': grpc.unary_unary_rpc_method_handler(
          servicer.ReserveStorage,
          request_deserializer=pool_dot_v0_dot_node__pb2.ReserveStorageRequest.FromString,
          response_serializer=pool_dot_v0_dot_node__pb2.Node.SerializeToString,
      ),
      'ReleaseStorage': grpc.unary_unary_rpc_method_handler(
          servicer.ReleaseStorage,
          request_deserializer=pool_dot_v0_dot_node__pb2.ReleaseStorageRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'n0stack.pool.NodeService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
