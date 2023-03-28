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
}
