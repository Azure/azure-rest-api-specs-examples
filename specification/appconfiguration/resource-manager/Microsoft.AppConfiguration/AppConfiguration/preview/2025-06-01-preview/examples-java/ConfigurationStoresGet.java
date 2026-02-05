
/**
 * Samples for ConfigurationStores GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresGet.json
     */
    /**
     * Sample code: ConfigurationStores_Get.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        configurationStoresGet(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().getByResourceGroupWithResponse("myResourceGroup", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
