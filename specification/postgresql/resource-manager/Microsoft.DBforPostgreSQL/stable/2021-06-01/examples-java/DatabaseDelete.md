Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresqlflexibleserver_1.0.0-beta.4/sdk/postgresqlflexibleserver/azure-resourcemanager-postgresqlflexibleserver/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Databases Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/DatabaseDelete.json
     */
    /**
     * Sample code: Delete a database.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void deleteADatabase(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.databases().delete("TestGroup", "testserver", "db1", Context.NONE);
    }
}
```
