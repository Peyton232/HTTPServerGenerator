
import 'package:dio/dio.dart';

void getPets() async {
  // try {
    var response = await Dio().get('http://127.0.0.1:42069');
    print(response);
  // } catch (e) {
  //   print(e);
  // }
}