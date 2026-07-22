import 'package:flutter_test/flutter_test.dart';
import 'package:pet_territory/config/app_environment.dart';

void main() {
  test('requires resolution before accessing current', () {
    expect(() => AppEnvironment.current, throwsA(isA<StateError>()));
  });

  group('AppEnvironment.fromValue', () {
    test('defaults to development', () {
      expect(
        AppEnvironment.fromValue(null).type,
        AppEnvironmentType.development,
      );
    });

    test('accepts explicit development', () {
      expect(
        AppEnvironment.fromValue('development').type,
        AppEnvironmentType.development,
      );
    });

    test('accepts explicit production', () {
      expect(
        AppEnvironment.fromValue('production').type,
        AppEnvironmentType.production,
      );
    });

    test('rejects unsupported values', () {
      expect(
        () => AppEnvironment.fromValue('staging'),
        throwsA(
          isA<ArgumentError>().having((error) => error.name, 'name', 'APP_ENV'),
        ),
      );
    });
  });
}
