
/**
 * Samples for ConfigurationStores PurgeDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/DeletedConfigurationStoresPurge.json
     */
    /**
     * Sample code: Purge a deleted configuration store.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        purgeADeletedConfigurationStore(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().purgeDeleted("westus", "contoso", com.azure.core.util.Context.NONE);
    }
}
