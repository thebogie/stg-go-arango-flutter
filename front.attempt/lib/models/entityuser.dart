import 'dart:ffi';

class EntityUser {
  final String id;
  final String fullname;
  final String email;
  final String password;
  final Bool active;

  EntityUser(this.id, this.fullname, this.email, this.password, this.active);
}
