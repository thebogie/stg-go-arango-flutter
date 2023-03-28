enum AppPage { home, login, register }

extension AppPageExtension on AppPage {
  String get toPath {
    switch (this) {
      case AppPage.login:
        return '/login';
      case AppPage.register:
        return '/register';
      case AppPage.home:
        return '/';
      default:
        return '/onboarding';
    }
  }

  String get toName {
    switch (this) {
      case AppPage.login:
        return "Login";
      case AppPage.register:
        return "Register";
      case AppPage.home:
        return "Home";
      default:
        return "Onboarding";
    }
  }
}
