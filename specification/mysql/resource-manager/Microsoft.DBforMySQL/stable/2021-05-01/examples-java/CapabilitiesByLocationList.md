Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysqlflexibleserver_1.0.0-beta.1/sdk/mysqlflexibleserver/azure-resourcemanager-mysqlflexibleserver/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for LocationBasedCapabilities List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/CapabilitiesByLocationList.json
     */
    /**
     * Sample code: CapabilitiesList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void capabilitiesList(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.locationBasedCapabilities().list("WestUS", Context.NONE);
    }
}
```
