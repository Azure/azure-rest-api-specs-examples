```java
import com.azure.core.util.Context;

/** Samples for ConfigurationStores PurgeDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/DeletedConfigurationStoresPurge.json
     */
    /**
     * Sample code: Purge a deleted configuration store.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void purgeADeletedConfigurationStore(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().purgeDeleted("westus", "contoso", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appconfiguration_1.0.0-beta.5/sdk/appconfiguration/azure-resourcemanager-appconfiguration/README.md) on how to add the SDK to your project and authenticate.
