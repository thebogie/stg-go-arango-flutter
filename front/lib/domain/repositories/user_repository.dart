// lib/domain/repositories/user_repository.dart

abstract class UserRepository {
  Future<String> login(String email, String password);
}
