Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appconfiguration_1.0.0-beta.4/sdk/appconfiguration/azure-resourcemanager-appconfiguration/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appconfiguration.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.appconfiguration.models.ConfigurationResourceType;

/** Samples for Operations RegionalCheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/RegionalCheckNameAvailable.json
     */
    /**
     * Sample code: ConfigurationStores_CheckNameAvailable.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCheckNameAvailable(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .operations()
            .regionalCheckNameAvailabilityWithResponse(
                "westus",
                new CheckNameAvailabilityParameters()
                    .withName("contoso")
                    .withType(ConfigurationResourceType.MICROSOFT_APP_CONFIGURATION_CONFIGURATION_STORES),
                Context.NONE);
    }
}
```
