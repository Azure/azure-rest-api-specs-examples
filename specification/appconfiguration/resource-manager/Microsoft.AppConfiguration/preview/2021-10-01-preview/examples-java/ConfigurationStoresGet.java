import com.azure.core.util.Context;

/** Samples for ConfigurationStores GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresGet.json
     */
    /**
     * Sample code: ConfigurationStores_Get.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresGet(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().getByResourceGroupWithResponse("myResourceGroup", "contoso", Context.NONE);
    }
}
