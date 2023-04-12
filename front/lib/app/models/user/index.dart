class User {
  final String email;
  final bool isLoggedIn;
  final String password;
  final String? accesstoken;
  User(
      {required this.email,
      this.password = "",
      this.isLoggedIn = false,
      this.accesstoken});

  factory User.fromJson(Map<String, dynamic> data) {
    final email = data['email'] as String; // cast as non-nullable String
    final password = data['password'] as String; // cast as non-nullable String
    final isLoggedIn = data['isLoggedIn'] as bool;
    final accesstoken = data['accesstoken'] as String?; // cast as nullable int
    return User(
        email: email,
        password: password,
        isLoggedIn: isLoggedIn,
        accesstoken: accesstoken ?? "");
  }

  Map<String, dynamic> toJson() => {
        "email": email,
        "password": password,
        "isLoggedIn": isLoggedIn,
        if (accesstoken != null) 'accesstoken': accesstoken,
      };
}
