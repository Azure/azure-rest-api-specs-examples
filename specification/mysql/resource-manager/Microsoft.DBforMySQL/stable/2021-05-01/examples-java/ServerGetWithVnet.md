Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysqlflexibleserver_1.0.0-beta.2/sdk/mysqlflexibleserver/azure-resourcemanager-mysqlflexibleserver/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Servers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerGetWithVnet.json
     */
    /**
     * Sample code: Get a server with vnet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getAServerWithVnet(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().getByResourceGroupWithResponse("testrg", "mysqltestserver", Context.NONE);
    }
}
```
