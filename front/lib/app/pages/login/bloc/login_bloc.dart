import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:smacktalkgaming/app/services/auth/index.dart';
import 'package:smacktalkgaming/app/services/api/index.dart' show ApiService;

import '../../../models/user/index.dart';
import 'package:bloc/bloc.dart';
import 'package:meta/meta.dart';

part 'login_event.dart';
part 'login_state.dart';

class LoginBloc extends Bloc<LoginEvent, LoginState> {
  final ApiService apiService;

  LoginBloc({required this.apiService}) : super(LoginInitial());

  @override
  Stream<LoginState> mapEventToState(LoginEvent event) async* {
    if (event is LoginButtonPressed) {
      yield LoginLoading();
      try {
        final user = await apiService.getUserFromAPI(
          event.email,
          event.password,
        );
        yield LoginSuccess(user: user);
      } catch (error) {
        yield LoginFailure(error: error.toString());
      }
    }
  }
}
