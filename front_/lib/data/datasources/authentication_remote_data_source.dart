// lib/data/datasources/authentication_remote_data_source.dart

import 'package:http/http.dart' as http;

class AuthenticationRemoteDataSource {
  static const apiUrl = 'https://your-api-url.com';

  Future<String> login(String email, String password) async {
    final url = '$apiUrl/login';
    final response = await http.post(Uri.parse(url), body: {
      'email': email,
      'password': password,
    });

    if (response.statusCode == 200) {
      final token = response.body;
      return token;
    } else {
      throw Exception('Login failed');
    }
  }
}
