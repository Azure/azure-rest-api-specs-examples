```java
import com.azure.core.util.Context;

/** Samples for ServerAdministrators List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerAdminList.json
     */
    /**
     * Sample code: get a list of server administrators.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getAListOfServerAdministrators(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.serverAdministrators().list("testrg", "mysqltestsvc4", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.
