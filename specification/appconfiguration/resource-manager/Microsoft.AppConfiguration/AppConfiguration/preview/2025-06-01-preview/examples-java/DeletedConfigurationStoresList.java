
/**
 * Samples for ConfigurationStores ListDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/DeletedConfigurationStoresList.json
     */
    /**
     * Sample code: DeletedConfigurationStores_List.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        deletedConfigurationStoresList(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().listDeleted(com.azure.core.util.Context.NONE);
    }
}
