Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets a private link resource for PostgreSQL.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getsAPrivateLinkResourceForPostgreSQL(
        com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.privateLinkResources().getWithResponse("Default", "test-svr", "plr", Context.NONE);
    }
}
```
