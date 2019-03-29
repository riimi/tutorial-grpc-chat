// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: gopher.proto
// </auto-generated>
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Protocol {
  public static partial class GopherService
  {
    static readonly string __ServiceName = "protocol.GopherService";

    static readonly grpc::Marshaller<global::Protocol.LoginRequest> __Marshaller_protocol_LoginRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Protocol.LoginRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Protocol.LoginResponse> __Marshaller_protocol_LoginResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Protocol.LoginResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Protocol.LogoutRequest> __Marshaller_protocol_LogoutRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Protocol.LogoutRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Protocol.LogoutResponse> __Marshaller_protocol_LogoutResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Protocol.LogoutResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Protocol.Message> __Marshaller_protocol_Message = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Protocol.Message.Parser.ParseFrom);

    static readonly grpc::Method<global::Protocol.LoginRequest, global::Protocol.LoginResponse> __Method_login = new grpc::Method<global::Protocol.LoginRequest, global::Protocol.LoginResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "login",
        __Marshaller_protocol_LoginRequest,
        __Marshaller_protocol_LoginResponse);

    static readonly grpc::Method<global::Protocol.LogoutRequest, global::Protocol.LogoutResponse> __Method_logout = new grpc::Method<global::Protocol.LogoutRequest, global::Protocol.LogoutResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "logout",
        __Marshaller_protocol_LogoutRequest,
        __Marshaller_protocol_LogoutResponse);

    static readonly grpc::Method<global::Protocol.Message, global::Protocol.Message> __Method_move = new grpc::Method<global::Protocol.Message, global::Protocol.Message>(
        grpc::MethodType.DuplexStreaming,
        __ServiceName,
        "move",
        __Marshaller_protocol_Message,
        __Marshaller_protocol_Message);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Protocol.GopherReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of GopherService</summary>
    public abstract partial class GopherServiceBase
    {
      public virtual global::System.Threading.Tasks.Task<global::Protocol.LoginResponse> login(global::Protocol.LoginRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Protocol.LogoutResponse> logout(global::Protocol.LogoutRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task move(grpc::IAsyncStreamReader<global::Protocol.Message> requestStream, grpc::IServerStreamWriter<global::Protocol.Message> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for GopherService</summary>
    public partial class GopherServiceClient : grpc::ClientBase<GopherServiceClient>
    {
      /// <summary>Creates a new client for GopherService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public GopherServiceClient(grpc::Channel channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for GopherService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public GopherServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected GopherServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected GopherServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::Protocol.LoginResponse login(global::Protocol.LoginRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return login(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Protocol.LoginResponse login(global::Protocol.LoginRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_login, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Protocol.LoginResponse> loginAsync(global::Protocol.LoginRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return loginAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Protocol.LoginResponse> loginAsync(global::Protocol.LoginRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_login, null, options, request);
      }
      public virtual global::Protocol.LogoutResponse logout(global::Protocol.LogoutRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return logout(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Protocol.LogoutResponse logout(global::Protocol.LogoutRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_logout, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Protocol.LogoutResponse> logoutAsync(global::Protocol.LogoutRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return logoutAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Protocol.LogoutResponse> logoutAsync(global::Protocol.LogoutRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_logout, null, options, request);
      }
      public virtual grpc::AsyncDuplexStreamingCall<global::Protocol.Message, global::Protocol.Message> move(grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return move(new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncDuplexStreamingCall<global::Protocol.Message, global::Protocol.Message> move(grpc::CallOptions options)
      {
        return CallInvoker.AsyncDuplexStreamingCall(__Method_move, null, options);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override GopherServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new GopherServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(GopherServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_login, serviceImpl.login)
          .AddMethod(__Method_logout, serviceImpl.logout)
          .AddMethod(__Method_move, serviceImpl.move).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the  service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static void BindService(grpc::ServiceBinderBase serviceBinder, GopherServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_login, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Protocol.LoginRequest, global::Protocol.LoginResponse>(serviceImpl.login));
      serviceBinder.AddMethod(__Method_logout, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Protocol.LogoutRequest, global::Protocol.LogoutResponse>(serviceImpl.logout));
      serviceBinder.AddMethod(__Method_move, serviceImpl == null ? null : new grpc::DuplexStreamingServerMethod<global::Protocol.Message, global::Protocol.Message>(serviceImpl.move));
    }

  }
}
#endregion