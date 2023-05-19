// lib/domain/usecases/login_usecase.dart

import 'package:smacktalkgaming/domain/repositories/user_repository.dart';

class LoginUseCase {
  final UserRepository userRepository;

  LoginUseCase(this.userRepository);

  Future<String> execute(String email, String password) async {
    return await userRepository.login(email, password);
  }
}
