enum AppEnvironmentType { development, production }

class AppEnvironment {
  const AppEnvironment._(this.type);

  static const _defaultValue = 'development';
  static AppEnvironment? _resolved;

  final AppEnvironmentType type;

  static AppEnvironment resolve() {
    return _resolved ??= fromValue(
      const String.fromEnvironment('APP_ENV', defaultValue: _defaultValue),
    );
  }

  static AppEnvironment get current {
    final environment = _resolved;
    if (environment == null) {
      throw StateError(
        'AppEnvironment.resolve() must be called before accessing AppEnvironment.current',
      );
    }

    return environment;
  }

  static AppEnvironment fromValue(String? value) {
    switch (value ?? _defaultValue) {
      case 'development':
        return const AppEnvironment._(AppEnvironmentType.development);
      case 'production':
        return const AppEnvironment._(AppEnvironmentType.production);
      default:
        throw ArgumentError.value(
          value,
          'APP_ENV',
          'must be development or production',
        );
    }
  }
}
