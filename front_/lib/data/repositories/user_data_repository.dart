// lib/data/repositories/user_data_repository.dart

import 'package:smacktalkgaming/data/datasources/authentication_remote_data_source.dart';
import 'package:smacktalkgaming/domain/repositories/user_repository.dart';

class UserDataRepository implements UserRepository{
  final AuthenticationRemoteDataSource authenticationRemoteDataSource;

  UserDataRepository(this.authenticationRemoteDataSource);

  Future<String> login(String email, String password) async {
    return await authenticationRemoteDataSource.login(email, password);
  }
}
