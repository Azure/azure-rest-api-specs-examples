
import com.azure.resourcemanager.appconfiguration.models.RegenerateKeyParameters;

/**
 * Samples for ConfigurationStores RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresRegenerateKey.json
     */
    /**
     * Sample code: ConfigurationStores_RegenerateKey.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        configurationStoresRegenerateKey(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().regenerateKeyWithResponse("myResourceGroup", "contoso",
            new RegenerateKeyParameters().withId("439AD01B4BE67DB1"), com.azure.core.util.Context.NONE);
    }
}
