import 'package:smacktalkgaming/app/models/user/index.dart';
import 'package:smacktalkgaming/app/services/auth/interface/auth_interface.dart'
    show AuthInterface;
import 'package:smacktalkgaming/app/services/api/index.dart' show ApiService;

class AuthService implements AuthInterface {
  final _apiService = ApiService();

  @override
  Future<User> performLogin() async {
    print("AuthService");
    await _apiService.getUserFromAPI("", "");
    print("End AuthService");
    return Future(() => User(email: "test@gmail.com", isLoggedIn: true));
  }
}
