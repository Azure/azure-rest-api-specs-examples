```java
import com.azure.core.util.Context;

/** Samples for ServerAdministrators Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerAdminDelete.json
     */
    /**
     * Sample code: ServerAdministratorsDelete.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverAdministratorsDelete(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.serverAdministrators().delete("testrg", "mysqltestsvc4", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.
