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


> **Pergunta:** Tente simular uma violação do pilar de Integridade.



> **Pergunta:** Tente simular uma violação do pilar de Disponibilidade. **Esse tem um truque!**

