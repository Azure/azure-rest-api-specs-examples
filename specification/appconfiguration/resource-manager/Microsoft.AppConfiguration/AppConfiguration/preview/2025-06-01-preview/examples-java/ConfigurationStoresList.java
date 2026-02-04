
/**
 * Samples for ConfigurationStores List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresList.json
     */
    /**
     * Sample code: ConfigurationStores_List.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        configurationStoresList(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().list(null, com.azure.core.util.Context.NONE);
    }
}
