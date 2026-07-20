import 'package:flutter/widgets.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:pet_territory/app/app.dart';

void main() {
  testWidgets('renders an empty app', (tester) async {
    await tester.pumpWidget(const PetTerritoryApp());

    expect(find.byType(SizedBox), findsOneWidget);
  });
}
