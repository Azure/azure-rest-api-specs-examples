Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.mysql.fluent.models.ConfigurationListResultInner;

/** Samples for ServerParameters ListUpdateConfigurations. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ConfigurationsUpdateByServer.json
     */
    /**
     * Sample code: ConfigurationList.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void configurationList(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .serverParameters()
            .listUpdateConfigurations("testrg", "mysqltestsvc1", new ConfigurationListResultInner(), Context.NONE);
    }
}
```
