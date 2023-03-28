import 'package:easy_localization/easy_localization.dart'
    show BuildContextEasyLocalizationExtension, StringTranslateExtension, tr;
import 'package:flutter/foundation.dart' show kDebugMode;
import 'package:flutter/material.dart'
    show
        AppBar,
        BuildContext,
        Center,
        MaterialApp,
        Scaffold,
        State,
        StatefulWidget,
        StatelessWidget,
        Text,
        TextStyle,
        Widget;
import 'package:flutter_bloc/flutter_bloc.dart'
    show BlocProvider, MultiBlocProvider;
import 'package:smacktalkgaming/app/config/routes/app_routes.dart'
    show AppRouter;
import 'package:smacktalkgaming/app/pages/auth/bloc/auth_bloc.dart'
    show AuthBloc;
import 'package:smacktalkgaming/app/services/auth/index.dart' show AuthService;

import 'config/theme/app_colors.dart' show AppColors;
import 'config/theme/app_theme.dart' show AppTheme;

final authBloc = AuthBloc(AuthService());

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [BlocProvider.value(value: authBloc)],
      child: MaterialApp.router(
        debugShowCheckedModeBanner: kDebugMode,
        routerConfig: AppRouter(authBloc).router,
        title: 'Smacktalk Gaming',
        theme: AppTheme.light,
        localizationsDelegates: [...context.localizationDelegates],
        supportedLocales: context.supportedLocales,
      ),
    );
  }
}

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(tr("Smacktalk Gaming")),
      ),
      body: Center(
        child: Text(
          "Smacktalk Gaming".tr(),
          style: TextStyle(color: AppColors.error),
        ),
      ),
    );
  }
}
