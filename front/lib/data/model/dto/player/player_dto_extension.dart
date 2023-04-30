import 'package:smacktalkgaming/data/model/dto/player/player_dto.dart';
import 'package:smacktalkgaming/data/model/remote/player/PlayerInfo.dart';

extension PlayerInfoExtension on PlayerInfo {
  PlayerDto toPlayerDto() {
    return PlayerDto(firstname, email, arangoid, "");
  }
}
