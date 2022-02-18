Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appconfiguration_1.0.0-beta.4/sdk/appconfiguration/azure-resourcemanager-appconfiguration/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ConfigurationStores ListDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/DeletedConfigurationStoresList.json
     */
    /**
     * Sample code: DeletedConfigurationStores_List.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void deletedConfigurationStoresList(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().listDeleted(Context.NONE);
    }
}
```
