
/**
 * Samples for ConfigurationStores GetDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/DeletedConfigurationStoresGet.json
     */
    /**
     * Sample code: DeletedConfigurationStores_Get.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        deletedConfigurationStoresGet(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().getDeletedWithResponse("westus", "contoso", com.azure.core.util.Context.NONE);
    }
}
