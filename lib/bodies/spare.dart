class Spare {
  final String type;
  final String number;
  final int weight;

  Spare(this.type, this.number, this.weight);

  Map<String, dynamic> toJson() {
    return {
      "Type": type,
      "Number": number,
      "Weight": weight,
    };
  }
}
