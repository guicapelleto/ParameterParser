Parameter Parser

Uma ferramenta de segurança ofensiva para ajudar na busca por subdomínios, comentários, parâmetros, arquivos JS, formulários entre outros elementos que podem ser explorados em aplicações web.
Possui a opção recursiva, onde busca em todas as páginas encontradas.

*Modo de uso*
paramparser [URL] --opcoes

Ex: paramparser https://meualvo.com -r -s

Opções:

OPÇÃO          ALIAS       DESCRIÇÃO
--recursivo    -r          Faz a verificação em todas as páginas recursivamente
--verbose      -v          Mostra todos os resultados na tela
--salvar       -s          Gera um relatório com os resultados obtidos
--help         -h          Mostra o menu de ajuda
