Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateEndpointConnectionList.json
     */
    /**
     * Sample code: Gets list of private endpoint connections on a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getsListOfPrivateEndpointConnectionsOnAServer(
        com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.privateEndpointConnections().listByServer("Default", "test-svr", Context.NONE);
    }
}
```
