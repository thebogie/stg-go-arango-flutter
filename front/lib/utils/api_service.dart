import 'dart:developer';

import 'package:get/get.dart';
import 'package:http/http.dart' as http;
//import 'package:rest_api_example/constants.dart';
//import 'package:rest_api_example/model/user_model.dart';

class ApiService {
  final _getConnect = GetConnect();

//  Future<List<UserModel>?> getUsers() async {
  Future<void> getUser(String email, String password) async {
    final response;

    try {
      response = await _getConnect.post(
        'http://localhost:50002/api/v1/login',
        {
          'Email': email,
          'Password': password,
        },
      );
      print(response);
      print('Response status: ${response.statusCode}');
      print('Response body: ${response.body}');

      //var response = await http.get(url);
      //if (response.statusCode == 200) {
      //  List<UserModel> _model = userModelFromJson(response.body);
      //  return _model;
      //}
    } catch (e) {
      log(e.toString());
    }
  }
}
