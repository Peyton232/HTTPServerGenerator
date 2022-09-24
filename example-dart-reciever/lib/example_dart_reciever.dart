import 'package:dio/dio.dart';

String host = "http://localhost:42069";

dynamic getPets() async {
  try {
    var response = await Dio().get(
      host + "/pets",
      // options: Options(
      //     //will not throw errors
      //     validateStatus: (status) => true,
      //   ),
      );
    return response.data;
  } catch (e) {
    return e.toString();
  }
}