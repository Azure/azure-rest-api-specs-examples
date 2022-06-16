import com.azure.core.util.Context;
import com.azure.resourcemanager.appconfiguration.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.appconfiguration.models.ConfigurationResourceType;

/** Samples for Operations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/CheckNameAvailable.json
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
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityParameters()
                    .withName("contoso")
                    .withType(ConfigurationResourceType.MICROSOFT_APP_CONFIGURATION_CONFIGURATION_STORES),
                Context.NONE);
    }
}
