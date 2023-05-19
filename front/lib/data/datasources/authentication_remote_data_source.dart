// lib/data/datasources/authentication_remote_data_source.dart

import 'package:dio/dio.dart';
import 'package:http/http.dart' as http;

class AuthenticationRemoteDataSource {
  static const apiUrl = 'http://127.0.0.1:50002/api/v1';

  Future<String> login(String email, String password) async {
    final url = '$apiUrl/login';
    final Dio _dio = Dio();
    try {
      final response = await _dio.post(url, data: {
        'Email': email,
        'Password': password,
      });

      if (response.statusCode == 200) {
        final token = response.data['data']['accessToken'];
        return token;
      } else {
        throw Exception('Login failed');
      }
    } catch (e) {
      throw Exception('Failed to connect to the server');
    }
  }
}
