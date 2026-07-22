import 'package:flutter/widgets.dart';

import 'app/app.dart';
import 'config/app_environment.dart';

void main() {
  AppEnvironment.resolve();
  runApp(const PetTerritoryApp());
}
