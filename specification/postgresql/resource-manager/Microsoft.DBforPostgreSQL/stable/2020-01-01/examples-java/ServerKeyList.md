```java
import com.azure.core.util.Context;

/** Samples for ServerKeys List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyList.json
     */
    /**
     * Sample code: List the keys for a PostgreSQL Server.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listTheKeysForAPostgreSQLServer(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverKeys().list("testrg", "testserver", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.
