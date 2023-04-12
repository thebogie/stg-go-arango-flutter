import '../../../models/user/index.dart';

abstract class ApiInterface {
  Future<User> getUserFromAPI(String email, String password);
  Future<User> loginUser(String email, String password);
}
