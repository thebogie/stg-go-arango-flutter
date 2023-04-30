class PlayerInfo {
  const PlayerInfo._({
    this.firstname,
    this.email,
    this.arangoid,
    this.password,
  });
  final String? firstname;
  final String? email;
  final String? arangoid;
  final String? password;

  static PlayerInfo fromJson(Map<String, dynamic> json) {
    return PlayerInfo._(
      firstname: json['firstname'],
      email: json['email'],
      arangoid: json['arangoid'],
      password: json['password'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'firstname': firstname,
      'email': email,
      'arangoid': arangoid,
      'password': password,
    };
  }
}
