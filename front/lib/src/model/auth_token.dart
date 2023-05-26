import 'dart:convert';

class AuthToken {
  AuthToken({
    this.code,
    this.message,
    this.data,
    this.token,
    this.userEmail,
    this.userNicename,
    this.userDisplayName,
  });

  String? code;
  String? message;
  Data? data;
  String? token;
  String? userEmail;
  String? userNicename;
  String? userDisplayName;

  factory AuthToken.fromJson(String str) => AuthToken.fromMap(json.decode(str));

  String toJson() => json.encode(toMap());

  factory AuthToken.fromMap(Map<String, dynamic> json) {
    print("JSON $json");
    return AuthToken(
      code: json["code"],
      message: json["message"],
      data: Data.fromMap(json["data"] ?? {}),
      token: json["data"]["accessToken"],
      userEmail: json["data"]["Email"],
      userNicename: json["data"]["Fullname"],
      userDisplayName: json["data"]["Fullname"],
    );
  }

  Map<String, dynamic> toMap() => {
        "code": code,
        "message": message,
        "data": data?.toMap(),
        "token": token,
        "user_email": userEmail,
        "user_nicename": userNicename,
        "user_display_name": userDisplayName,
      };
  @override
  String toString() {
    JsonEncoder encoder = const JsonEncoder.withIndent('    ');
    return encoder.convert(toMap());
  }
}

class Data {
  Data({
    this.status,
  });

  int? status;

  factory Data.fromJson(String str) => Data.fromMap(json.decode(str));

  String toJson() => json.encode(toMap());

  factory Data.fromMap(Map<String, dynamic> json) => Data(
        status: json["status"],
      );

  Map<String, dynamic> toMap() => {
        "status": status,
      };
}
