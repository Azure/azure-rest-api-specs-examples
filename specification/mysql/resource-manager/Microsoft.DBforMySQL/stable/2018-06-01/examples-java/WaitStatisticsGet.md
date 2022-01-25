Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WaitStatistics Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/WaitStatisticsGet.json
     */
    /**
     * Sample code: WaitStatisticsGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void waitStatisticsGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .waitStatistics()
            .getWithResponse(
                "testResourceGroupName",
                "testServerName",
                "636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0",
                Context.NONE);
    }
}
```
