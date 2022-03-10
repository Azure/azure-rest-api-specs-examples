Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysqlflexibleserver_1.0.0-beta.2/sdk/mysqlflexibleserver/azure-resourcemanager-mysqlflexibleserver/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Servers Failover. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ServerFailover.json
     */
    /**
     * Sample code: Restart a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void restartAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.servers().failover("TestGroup", "testserver", Context.NONE);
    }
}
```
