Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Servers Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerStop.json
     */
    /**
     * Sample code: ServerStop.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void serverStop(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.servers().stop("TestGroup", "testserver", Context.NONE);
    }
}
```
