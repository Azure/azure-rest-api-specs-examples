/** Samples for ConfigurationStores GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresGet.json
     */
    /**
     * Sample code: ConfigurationStores_Get.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresGet(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .configurationStores()
            .getByResourceGroupWithResponse("myResourceGroup", "contoso", com.azure.core.util.Context.NONE);
    }
}
