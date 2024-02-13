# Ports and Adapters

Separação de lógica de negócio e tecnologias que serão usadas.

# Important things about architecture 

- Crescimento sustentável (precisa deixar que novas coisas possam ser adicionadas)
- Software precisa se pagar ao passar do tempo
- Software deve ser desenhado por você e não pelo seu framework
- Peças precisam se encaixar e evetualmente podem ser substituídas


# Main Problems

- Visão de futuro
- Limites bem definidos
- Troca e adição de componentes
- Escala
- Otimização frequentes
- Preparado para mudanças

# DESIGN Vs ARCHITECTURE

O design é algo micro da aplicação(ex: solid), já a arquitetura engloba todo o projeto. Uma mudança de classe é uma mudança de design, já uma mudança de ferramentas do projeto é uma mudança de arquitetura.


### Ports and Adapters

- Definição de limites e proteção nas regras da aplicação
- Componentização e desacoplamento
    - Logs
    - Cache
    - Upload
    - Banco de Dados
    - Comandos
    - Filas
    - Http / Apis / GraphQL
- Facilidade na quebra para microsserviços
- Não há padrão estabelecido de como o código deve ser organizado(ex: DDD)


Lado esquerdo do hexagono = coisas relacionadas ao cliente, exemplo: quem vai acessar nossa aplicação(HTTP, CLI)
Lado direito do hexagono = coisas relacionadas ao servidor, exemplo: conexões com banco de dados, filas, logs.

### SOLID with ports and adapters

- D = Dependency Inversion Principle
    - Módulos de alto nível não devem depender de módulos de baixo nível. Ambos devem depender de abstrações
    - Abstrações não devem depender de detalhes. Detalhes devem depender de abstrações