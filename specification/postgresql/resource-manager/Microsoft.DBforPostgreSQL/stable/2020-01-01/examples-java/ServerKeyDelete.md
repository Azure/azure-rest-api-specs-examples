```java
import com.azure.core.util.Context;

/** Samples for ServerKeys Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyDelete.json
     */
    /**
     * Sample code: Delete the PostgreSQL Server key.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void deleteThePostgreSQLServerKey(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .serverKeys()
            .delete("testserver", "someVault_someKey_01234567890123456789012345678901", "testrg", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.
