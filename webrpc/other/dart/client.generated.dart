// hello-world v0.1.0 4b39792dcfb717dadcad792cc6d0bad614978b92
// --
// Code generated by webrpc-gen@v0.19.3 with dart generator. DO NOT EDIT.
//
// webrpc-gen -schema=schema.ridl -target=dart -client -out=other/dart/client.generated.dart
// ignore_for_file: non_constant_identifier_names
// ignore_for_file: constant_identifier_names
import 'dart:convert';
import 'dart:async';
import 'package:http/http.dart' as http;

/// WebRPC description and code-gen version
const webRPCVersion = "v1";

/// Schema version of your RIDL schema
const webRPCSchemaVersion = "v0.1.0";

/// Schema hash generated from your RIDL schema
const webRPCSchemaHash = "4b39792dcfb717dadcad792cc6d0bad614978b92";

class HelloWorldV1Impl implements HelloWorldV1 {
    HelloWorldV1Impl(String hostname, [WebrpcHttpClient? httpClient])
        : _baseUrl = '$hostname/rpc/HelloWorldV1/',
          _httpClient = httpClient ?? _MainWebrpcHttpClient();

    final String _baseUrl;
    final WebrpcHttpClient _httpClient;

    @override
    Future<({String message})> simple() async {
        final String? body = null;
        WebrpcHttpRequest request = WebrpcHttpRequest(
            uri: _makeUri('Simple'),
            headers: _makeHeaders(),
            body: body,
        );
        final WebrpcHttpResponse response = await _httpClient.post(request);
        
        await _handleResponse(response);
        final Map<String, dynamic> json = jsonDecode(response.body);
        return (
            message: _simpleMessage(json['message']),
        );
    }
    
    static String _simpleMessage(dynamic v0) {
        final String r0 = _cast<String>(v0);
        return r0;
    }

    @override
    Future<({String message})> greet(User user) async {
        final String body = jsonEncode(toJsonObject({
            'user': user,
        }));
        WebrpcHttpRequest request = WebrpcHttpRequest(
            uri: _makeUri('Greet'),
            headers: _makeHeaders(),
            body: body,
        );
        final WebrpcHttpResponse response = await _httpClient.post(request);
        
        await _handleResponse(response);
        final Map<String, dynamic> json = jsonDecode(response.body);
        return (
            message: _greetMessage(json['message']),
        );
    }
    
    static String _greetMessage(dynamic v0) {
        final String r0 = _cast<String>(v0);
        return r0;
    }

    Uri _makeUri(String name) {
        return Uri.parse(_baseUrl + name);
    }

    static Map<String, Object> _makeHeaders() {
        return {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        };
    }

    Future<void> _handleResponse(WebrpcHttpResponse response) async {
        if (response.statusCode >= 400) {
            try {
                final Map<String, dynamic> json = jsonDecode(response.body);
                final int webrpcErrorCode = json['code'];
                if (response.statusCode >= 500) {
                    throw WebrpcException.fromCode(webrpcErrorCode);

                } else {
                    throw WebrpcError.fromCode(webrpcErrorCode);
                }
            } on ArgumentError catch (_) {
                // https://github.com/webrpc/webrpc/blob/master/gen/errors.go
                throw WebrpcException.fromCode(-5);
            }
        }
    }
}

class _MainWebrpcHttpClient implements WebrpcHttpClient {
  @override
  Future<WebrpcHttpResponse> post(WebrpcHttpRequest request) async {
    final http.Response response = await http.post(
      request.uri,
      body: request.body == null ? null : utf8.encode(request.body!),
      headers: request.headers.map((key, value) => MapEntry(key, value.toString())),
    );
    return WebrpcHttpResponse(statusCode: response.statusCode, body: response.body);
  }
}

abstract interface class WebrpcHttpClient {
  Future<WebrpcHttpResponse> post(WebrpcHttpRequest request);
}

class WebrpcHttpRequest {
  WebrpcHttpRequest({
    required this.uri,
    required this.headers,
    required this.body,
  });

  final Uri uri;
  final Map<String, Object> headers;
  final String? body;
}

class WebrpcHttpResponse {
  WebrpcHttpResponse({required this.statusCode, required this.body});

  final int statusCode;
  final String body;
}


enum Gender {
    MALE,
    FEMALE,
    OTHER;

    factory Gender.fromJson(dynamic json) {
        switch (json) {
            case 'MALE':
                return Gender.MALE;
            case 'FEMALE':
                return Gender.FEMALE;
            case 'OTHER':
                return Gender.OTHER;
            default:
                throw ArgumentError.value(json);
        }
    }

    String toJson() {
        return name;
    }
}

class Address implements JsonSerializable {
    Address({
        required this.street,
        required this.postalCode,
        required this.city,
        required this.country
    });
    
    final String street;
    final String? postalCode;
    final String city;
    final String country;
  
    Address.fromJson(Map<String, dynamic> json)
        : street = _street(json['street']),
          postalCode = _postalCode(json['postalCode']),
          city = _city(json['city']),
          country = _country(json['country']);
          
    static String _street(dynamic v0) {
        if (v0 == null) throw WebrpcException.fromCode(ErrorId.webrpcBadResponse.code);
        final String r0 = _cast<String>(v0);
        return r0;
    }
    
    static String? _postalCode(dynamic v0) {
        if (v0 == null) return null;
        final String? r0 = _cast<String?>(v0);
        return r0;
    }
    
    static String _city(dynamic v0) {
        if (v0 == null) throw WebrpcException.fromCode(ErrorId.webrpcBadResponse.code);
        final String r0 = _cast<String>(v0);
        return r0;
    }
    
    static String _country(dynamic v0) {
        if (v0 == null) throw WebrpcException.fromCode(ErrorId.webrpcBadResponse.code);
        final String r0 = _cast<String>(v0);
        return r0;
    }
    
    @override
    Map<String, dynamic> toJson() {
        return {
            'street': toJsonObject(street),
            'postalCode': toJsonObject(postalCode),
            'city': toJsonObject(city),
            'country': toJsonObject(country),
        };
    }
}

class User implements JsonSerializable {
    User({
        required this.name,
        required this.age,
        required this.gender,
        required this.address
    });
    
    final String name;
    final int age;
    final Gender gender;
    final Address? address;
  
    User.fromJson(Map<String, dynamic> json)
        : name = _name(json['name']),
          age = _age(json['age']),
          gender = _gender(json['gender']),
          address = _address(json['address']);
          
    static String _name(dynamic v0) {
        if (v0 == null) throw WebrpcException.fromCode(ErrorId.webrpcBadResponse.code);
        final String r0 = _cast<String>(v0);
        return r0;
    }
    
    static int _age(dynamic v0) {
        if (v0 == null) throw WebrpcException.fromCode(ErrorId.webrpcBadResponse.code);
        final int r0 = _cast<int>(v0);
        return r0;
    }
    
    static Gender _gender(dynamic v0) {
        if (v0 == null) throw WebrpcException.fromCode(ErrorId.webrpcBadResponse.code);
        final Gender r0 = Gender.fromJson(v0);
        return r0;
    }
    
    static Address? _address(dynamic v0) {
        if (v0 == null) return null;
        final Address r0 = Address.fromJson(v0);
        return r0;
    }
    
    @override
    Map<String, dynamic> toJson() {
        return {
            'name': toJsonObject(name),
            'age': toJsonObject(age),
            'gender': toJsonObject(gender),
            'address': toJsonObject(address),
        };
    }
}


T _cast<T>(x) {
    if ((x == null) && (null is T)) {
        return x;
    } else if (x is T) {
        return x;
    } else {
        throw ArgumentError.value(x);
    }
}

dynamic toJsonObject(dynamic v) {
    if (v == null) return null;
    if (v is DateTime) return v.toIso8601String();
    if (v is BigInt) return v.toString();
    // records are impossible to JSON serialize accurately because they do not
    // retain runtime info about their structure
    // see https://github.com/dart-lang/language/issues/2826
    if (v is Record) return v.toString();
    if (v is List) return v.map(toJsonObject).toList();
    if (v is Map) return v.map((key, value) => MapEntry(key.toString(), toJsonObject(value)));
    if (v is JsonSerializable) return v.toJson();
    return v;
}

DateTime? _dateTimeFromJsonOptional(dynamic v0) {
    if (v0 == null) return null;
    return _dateTimeFromJson(v0);
}

DateTime _dateTimeFromJson(dynamic v0) {
    if ((v0 != null) && (v0 is String)) {
        return DateTime.parse(v0);
    } else {
        throw ArgumentError.value(v0, "v0", "Cannot parse to DateTime");
    }
}

BigInt? _bigIntFromJsonOptional(dynamic v0) {
    if (v0 == null) return null;
    return _bigIntFromJson(v0);
}

BigInt _bigIntFromJson(dynamic v0) {
    if (v0 is String) {
        return BigInt.parse(v0);
    } else if (v0 is int) {
        return BigInt.from(v0);
    } else {
        throw ArgumentError.value(v0, "v0", "Required non-null BigInt");
    }
}

abstract interface class JsonSerializable {
    dynamic toJson();
}


abstract interface class HelloWorldV1 {
    Future<({String message})> simple();
    Future<({String message})> greet(User user);
}

/// Unrecoverable errors representing an invalid use of the API, bad schema, or
/// failure in the core of Webrpc (i.e. a bug).
class WebrpcError extends Error {
    WebrpcError._({
        required this.id, 
        required this.message, 
        required this.httpStatus,
    });
    
    factory WebrpcError.fromCode(int code) {
        switch (code) {
            case 0:
                return WebrpcError._(
                    id: ErrorId.webrpcEndpoint, 
                    message: 'endpoint error',
                    httpStatus: 400,
                );
            
            case -1:
                return WebrpcError._(
                    id: ErrorId.webrpcRequestFailed, 
                    message: 'request failed',
                    httpStatus: 400,
                );
            
            case -2:
                return WebrpcError._(
                    id: ErrorId.webrpcBadRoute, 
                    message: 'bad route',
                    httpStatus: 404,
                );
            
            case -3:
                return WebrpcError._(
                    id: ErrorId.webrpcBadMethod, 
                    message: 'bad method',
                    httpStatus: 405,
                );
            
            case -4:
                return WebrpcError._(
                    id: ErrorId.webrpcBadRequest, 
                    message: 'bad request',
                    httpStatus: 400,
                );
            
            case -8:
                return WebrpcError._(
                    id: ErrorId.webrpcClientDisconnected, 
                    message: 'client disconnected',
                    httpStatus: 400,
                );
            
            case -9:
                return WebrpcError._(
                    id: ErrorId.webrpcStreamLost, 
                    message: 'stream lost',
                    httpStatus: 400,
                );
            
            case -10:
                return WebrpcError._(
                    id: ErrorId.webrpcStreamFinished, 
                    message: 'stream finished',
                    httpStatus: 200,
                );
            
            default:
                throw ArgumentError.value(code, "code", "Unrecognized");
        }
    }

    final ErrorId id;
    final String message;
    final int httpStatus;
}

/// Recoverable errors that should generally be caught, representing a
/// bad state or temporary failure.
class WebrpcException implements Exception {
    WebrpcException._({
        required this.id, 
        required this.message, 
        required this.httpStatus,
    });
    
    factory WebrpcException.fromCode(int code) {
        switch (code) {
            case -5:
                return WebrpcException._(
                    id: ErrorId.webrpcBadResponse, 
                    message: 'bad response',
                    httpStatus: 500,
                );
            
            case -6:
                return WebrpcException._(
                    id: ErrorId.webrpcServerPanic, 
                    message: 'server panic',
                    httpStatus: 500,
                );
            
            case -7:
                return WebrpcException._(
                    id: ErrorId.webrpcInternalError, 
                    message: 'internal error',
                    httpStatus: 500,
                );
            
            default:
                throw ArgumentError.value(code, "code", "Unrecognized code $code");
        }
    }

    final ErrorId id;
    final String message;
    final int httpStatus;
}

/// Unique ID of a custom schema error or base Webrpc error.
enum ErrorId {
    webrpcEndpoint(code: 0, name: 'WebrpcEndpoint'),
    webrpcRequestFailed(code: -1, name: 'WebrpcRequestFailed'),
    webrpcBadRoute(code: -2, name: 'WebrpcBadRoute'),
    webrpcBadMethod(code: -3, name: 'WebrpcBadMethod'),
    webrpcBadRequest(code: -4, name: 'WebrpcBadRequest'),
    webrpcBadResponse(code: -5, name: 'WebrpcBadResponse'),
    webrpcServerPanic(code: -6, name: 'WebrpcServerPanic'),
    webrpcInternalError(code: -7, name: 'WebrpcInternalError'),
    webrpcClientDisconnected(code: -8, name: 'WebrpcClientDisconnected'),
    webrpcStreamLost(code: -9, name: 'WebrpcStreamLost'),
    webrpcStreamFinished(code: -10, name: 'WebrpcStreamFinished');

    const ErrorId({
        required this.code,
        required this.name,
    });

    final int code;
    final String name;
}

