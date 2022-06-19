import com.azure.core.util.Context;
import com.azure.resourcemanager.appconfiguration.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.appconfiguration.models.ConfigurationResourceType;

/** Samples for Operations CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/CheckNameNotAvailable.json
     */
    /**
     * Sample code: ConfigurationStores_CheckNameNotAvailable.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresCheckNameNotAvailable(
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
