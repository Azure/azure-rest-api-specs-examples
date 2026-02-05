
import com.azure.resourcemanager.appconfiguration.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.appconfiguration.models.ConfigurationResourceType;

/**
 * Samples for Operations RegionalCheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/RegionalCheckNameNotAvailable.json
     */
    /**
     * Sample code: ConfigurationStores_CheckNameNotAvailable.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCheckNameNotAvailable(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.operations().regionalCheckNameAvailabilityWithResponse("westus",
            new CheckNameAvailabilityParameters().withName("contoso")
                .withType(ConfigurationResourceType.MICROSOFT_APP_CONFIGURATION_CONFIGURATION_STORES),
            com.azure.core.util.Context.NONE);
    }
}
