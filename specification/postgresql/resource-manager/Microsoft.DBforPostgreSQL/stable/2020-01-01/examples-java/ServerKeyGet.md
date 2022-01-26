Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ServerKeys Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyGet.json
     */
    /**
     * Sample code: Get the PostgreSQL Server key.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getThePostgreSQLServerKey(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager
            .serverKeys()
            .getWithResponse(
                "testrg", "testserver", "someVault_someKey_01234567890123456789012345678901", Context.NONE);
    }
}
```
