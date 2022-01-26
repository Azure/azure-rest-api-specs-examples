Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateEndpointConnectionDelete.json
     */
    /**
     * Sample code: Deletes a private endpoint connection with a given name.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void deletesAPrivateEndpointConnectionWithAGivenName(
        com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .privateEndpointConnections()
            .delete("Default", "test-svr", "private-endpoint-connection-name", Context.NONE);
    }
}
```
