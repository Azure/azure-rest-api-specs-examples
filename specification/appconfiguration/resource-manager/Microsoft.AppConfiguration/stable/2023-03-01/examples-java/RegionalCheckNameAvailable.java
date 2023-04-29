import com.azure.resourcemanager.appconfiguration.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.appconfiguration.models.ConfigurationResourceType;

/** Samples for Operations RegionalCheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/RegionalCheckNameAvailable.json
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
                com.azure.core.util.Context.NONE);
    }
}
