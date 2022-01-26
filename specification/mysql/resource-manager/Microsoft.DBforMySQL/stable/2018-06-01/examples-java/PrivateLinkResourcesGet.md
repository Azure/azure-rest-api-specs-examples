Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets a private link resource for MySQL.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void getsAPrivateLinkResourceForMySQL(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager.privateLinkResources().getWithResponse("Default", "test-svr", "plr", Context.NONE);
    }
}
```
