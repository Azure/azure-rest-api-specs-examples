
/**
 * Samples for ConfigurationStores Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresDelete.json
     */
    /**
     * Sample code: ConfigurationStores_Delete.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        configurationStoresDelete(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().delete("myResourceGroup", "contoso", com.azure.core.util.Context.NONE);
    }
}
