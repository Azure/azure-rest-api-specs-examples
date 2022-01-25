Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ServerKeys Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerKeyDelete.json
     */
    /**
     * Sample code: Delete the MySQL Server key.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void deleteTheMySQLServerKey(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .serverKeys()
            .delete("testserver", "someVault_someKey_01234567890123456789012345678901", "testrg", Context.NONE);
    }
}
```
