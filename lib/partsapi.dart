import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/spare.dart';

Future<HttpRequest> createSpare(Spare obj) async {
  var apiroute = getEndpoint("parts");
  var url = "${apiroute}/info";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateSpare(Key key, Spare obj) async {
  var route = getEndpoint("parts");
  var url = "${route}/info/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteSpare(Key key) async {
  var route = getEndpoint("parts");
  var url = "${route}/info/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
