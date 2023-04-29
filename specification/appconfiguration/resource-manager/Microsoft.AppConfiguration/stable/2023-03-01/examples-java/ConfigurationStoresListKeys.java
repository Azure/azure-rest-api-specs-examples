/** Samples for ConfigurationStores ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresListKeys.json
     */
    /**
     * Sample code: ConfigurationStores_ListKeys.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresListKeys(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().listKeys("myResourceGroup", "contoso", null, com.azure.core.util.Context.NONE);
    }
}
