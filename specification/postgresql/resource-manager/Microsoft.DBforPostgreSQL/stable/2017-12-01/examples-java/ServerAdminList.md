```java
import com.azure.core.util.Context;

/** Samples for ServerAdministrators List. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerAdminList.json
     */
    /**
     * Sample code: get a list of server administrators.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getAListOfServerAdministrators(com.azure.resourcemanager.postgresql.PostgreSqlManager manager) {
        manager.serverAdministrators().list("testrg", "pgtestsvc4", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-postgresql_1.0.2/sdk/postgresql/azure-resourcemanager-postgresql/README.md) on how to add the SDK to your project and authenticate.
