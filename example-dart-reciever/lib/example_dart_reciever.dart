
import 'package:dio/dio.dart';

void getPets() async {
  // try {
    var response = await Dio().get('http://localhost:42069/pets');
    print(response);
  // } catch (e) {
  //   print(e);
  // }
}