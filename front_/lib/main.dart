// lib/main.dart

import 'package:flutter/material.dart';
import 'package:smacktalkgaming/data/datasources/authentication_remote_data_source.dart';
import 'package:smacktalkgaming/data/repositories/user_data_repository.dart';
import 'package:smacktalkgaming/domain/repositories/user_repository.dart';
import 'package:smacktalkgaming/domain/usecases/login_usecase.dart';
import 'package:smacktalkgaming/presentation/screens/login_screen.dart';

void main() {
  final AuthenticationRemoteDataSource authenticationRemoteDataSource =
      AuthenticationRemoteDataSource();
  final UserRepository userRepository =
      UserDataRepository(authenticationRemoteDataSource);
  final LoginUseCase loginUseCase = LoginUseCase(userRepository);

  runApp(MyApp(loginUseCase));
}

class MyApp extends StatelessWidget {
  final LoginUseCase loginUseCase;

  MyApp(this.loginUseCase);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Login Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: LoginScreen(loginUseCase),
    );
  }
}
