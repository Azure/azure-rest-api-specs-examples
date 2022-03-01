Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appconfiguration_1.0.0-beta.5/sdk/appconfiguration/azure-resourcemanager-appconfiguration/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appconfiguration.models.RegenerateKeyParameters;

/** Samples for ConfigurationStores RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresRegenerateKey.json
     */
    /**
     * Sample code: ConfigurationStores_RegenerateKey.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresRegenerateKey(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .configurationStores()
            .regenerateKeyWithResponse(
                "myResourceGroup", "contoso", new RegenerateKeyParameters().withId("439AD01B4BE67DB1"), Context.NONE);
    }
}
```
