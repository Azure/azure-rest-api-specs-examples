
import com.azure.resourcemanager.appconfiguration.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.appconfiguration.models.ConfigurationResourceType;

/**
 * Samples for Operations CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/CheckNameAvailable.json
     */
    /**
     * Sample code: ConfigurationStores_CheckNameAvailable.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCheckNameAvailable(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.operations().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityParameters().withName("contoso")
                .withType(ConfigurationResourceType.MICROSOFT_APP_CONFIGURATION_CONFIGURATION_STORES),
            com.azure.core.util.Context.NONE);
    }
}
