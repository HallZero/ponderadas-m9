# Simulação de ataques usando MQTT

Este arquivo possui as perguntas para o README, bem como a execução do passo-a-passo das instruções fornecidas.

[Broker remoto ClientID](https://youtu.be/B1fG-MlkN-E)

[MQTT sem autenticação](https://youtu.be/n3EICX47FLE)

[MQTT com autenticação](https://youtu.be/_WaSBiV0XZA)



> **Pergunta:** O que acontece se você utilizar o mesmo ClientID em outra máquina ou sessão do browser? Algum pilar do CIA Triad é violado com isso?

Ao utilizar o ClientID em outra máquina ou sessão, a primeira conexão irá se desconectar. Pensando em um cenário em que essa informação é vazada, o pilar de Confidencialidade é violado, uma vez que esta informação deveria ser única.

> **Pergunta:** Com os parâmetros de resources, algum pilar do CIA Triad pode ser facilmente violado?

Observando os materiais fornecidos, pode-se inferir violações envolvendo o pilar de Disponibilidade (Availability). Nesse contexto, a definição do uso de recursos computacionais definida neste nível pode afetar a performace por limites operacionais restritos.

> **Pergunta:** Sem autenticação (repare que a variável allow_anonymous está como true), como a parte de confidencialidade pode ser violada?

No contexto de segurança, pensando que qualquer usuário pode publicar informações e subscrever-se nos tópicos, a informação não é restrita apenas à pessoas autorizadas, que por definição já viola o conceito de confidencialidade.

> **Pergunta:** Tente simular uma violação do pilar de Confidencialidade.

O Pilar de Confidencialidade no Triad CIA é referente ao critério de que apenas pessoas autorizadas possuam acesso à informações. Pensando nesse contexto, a interceptação dessa mensagem entre o publisher e o broker ou entre o broker e subscriber poderia ocorrer por um atacante utilizando ferramentas de captura de pacotes enquato eles são transmitidos pela rede (local ou internet), visto que no exemplo acima, não existe proteção ou criptografia (SSL/TLS).

Imaginando um cenário hipotético em que utiliza-se o envio de dados de sensores localizados em uma região de uma metrópole para o monitoramento da saúde pública através da internet entre um publisher e um broker, poderiamos utilizar ferramentas populares como o **Wireshark** ou **tcpdump** para interceptar e analisar as informações de um payload que não é protegido com criptografia. Isso nos daria acesso ao conteúdos de mensagens que não necessariamente deveriam ser acessadas por terceiros sem autorização.

> **Pergunta:** Tente simular uma violação do pilar de Integridade.

Tendo em vista que o pilar de Integridade diz respeito à "Imacularidade" dos dados, isto é, a pureza no sentido de não violação do seu conteúdo, pode-se pensar em formas de infringir esse aspecto.

Utilizando o mesmo contexto da resposta anterior, uma técnica que poderia ser utilizada para violar a Integridade da mensagem seria a estratégia de *Man-In-The-Middle (MitM)*, que justamente visa interceptar e modificar uma mensagem entre as duas partes legítimas. Poderiamos então capturar uma mensagem enviada pelo publisher e modificar seu payload com dados falsos utilizando **ettercap**.


> **Pergunta:** Tente simular uma violação do pilar de Disponibilidade. **Esse tem um truque!**

Mantendo a ideia do contexto anterior, uma possibilidade de ideia seria indisponibilizar o serviço a partir de uma sobrecarga forçada do sistema num ponto de falha crítico (nesse caso, o broker), configurando um *Denial of Service (DoS)*. Nessa estratégia, um número excessivo de publishers poderia ser criado de forma a afetar a disponibilidade do Broker ao publicar mais mensagens do que pode-se lidar. Poderiamos utilizar o **LOIC (Low Orbit Ion Cannon)** para coordenar um ataque, embora brokers MQTT sejam projetados para lidar com um grande número de dados e conexões.

