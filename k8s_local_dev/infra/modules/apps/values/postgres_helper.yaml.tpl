extraEnv:
  - name: CONSUL_IP
    value: ${consul_ip_val}
  - name: VAULT_IP
    value: ${vault_ip_val}

initdbScripts:
  db-init.sql: |
    CREATE TABLE ValueToChange (
        val_name varchar(255),
        val int
    );

    INSERT INTO ValueToChange (val_name, val)
    VALUES ('caas_value', 40);