import 'dart:convert';

import 'package:deep_pick/deep_pick.dart';
import 'package:dio/dio.dart';
import 'package:smacktalkgaming/app/models/user/index.dart';
import 'package:smacktalkgaming/app/services/api/interface/api_interface.dart'
    show ApiInterface;
import 'package:smacktalkgaming/config.dart';

class ApiService implements ApiInterface {
  final dio = Dio();
  late final Response response;

  @override
  Future<User> getUserFromAPI(String email, String password) async {
    try {
      print("Hitting: ${Config.apiHost}/login");
      final response = await dio.get('${Config.apiHost}/login');
      print("Got: ${response.data.toString()}");
    } catch (e) {
      print("E: ${e.toString()}");
    }
    return Future(() => User(email: "test@gmail.com", isLoggedIn: true));
  }

  @override
  Future<User> loginUser(String email, String password) async {
    Map<String, dynamic> map = Map<String, dynamic>();

    try {
      print("Hitting: ${Config.apiHost}/login");
      final response = await dio.post('${Config.apiHost}/login',
          data: {'Email': email, 'Password': password},
          options: Options(
            responseType: ResponseType.plain,
          ));

      print("Got: ${response.data.toString()}");
      final json = jsonDecode(response.data)['data'];
      final String? accesstoken = pick(json, 'accessToken').asStringOrNull();

      map = {
        "accesstoken": pick(json, 'accessToken').asStringOrNull(),
        "email": pick(json, 'Email').asStringOrNull(),
        "password": pick(json, 'Password').asStringOrNull(),
        "isLoggedIn": false
      };

      User.fromJson(map);
    } catch (e) {
      print("E: ${e.toString()}");
    }
    return Future(() => User.fromJson(map));
  }
}
