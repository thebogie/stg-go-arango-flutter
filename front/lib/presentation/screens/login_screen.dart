// lib/presentation/screens/login_screen.dart

import 'package:flutter/material.dart';
import 'package:smacktalkgaming/domain/usecases/login_usecase.dart';

class LoginScreen extends StatelessWidget {
  final LoginUseCase loginUseCase;

  LoginScreen(this.loginUseCase);

  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();

  Future<void> _login(BuildContext context) async {
    final email = emailController.text;
    final password = passwordController.text;

    try {
      final token = await loginUseCase.execute(email, password);
      // Handle successful login, save token, navigate to the next screen, etc.
      print('Token: $token');
    } catch (e) {
      // Handle login failure, show error message, etc.
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Login'),
      ),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextField(
              controller: emailController,
              decoration: InputDecoration(labelText: 'Email'),
            ),
            TextField(
              controller: passwordController,
              decoration: InputDecoration(labelText: 'Password'),
              obscureText: true,
            ),
            SizedBox(height: 16.0),
            ElevatedButton(
              onPressed: () => _login(context),
              child: Text('Login'),
            ),
          ],
        ),
      ),
    );
  }
}
