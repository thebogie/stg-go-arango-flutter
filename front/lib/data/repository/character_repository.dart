import 'package:retrofit/retrofit.dart';
import 'package:smacktalkgaming/data/model/remote/character/CharacterInfo.dart';
import 'package:smacktalkgaming/data/model/remote/character/CharacterResponse.dart';
import 'package:smacktalkgaming/data/remote/service/character/character_service.dart';

class CharacterRepository {
  final CharacterService _characterService;

  CharacterRepository(this._characterService);

  Future<HttpResponse<CharacterResponse>> getCharacters(
    int page,
    Map<String, String>? options,
  ) =>
      _characterService.getCharacters(page, options);

  Future<HttpResponse<CharacterInfo>> getCharacter(int id) =>
      _characterService.getCharacter(id);
}
