class Config {
  static const apiUrl = "url";
  static const apiHost = String.fromEnvironment("API_HOST",
      defaultValue: "http://127.0.0.1:50002/api/v1");
  static const locale = "en";
  static const cacheDays = 1;
}
