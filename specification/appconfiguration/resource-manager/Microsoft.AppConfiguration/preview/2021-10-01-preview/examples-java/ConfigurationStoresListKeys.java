import com.azure.core.util.Context;

/** Samples for ConfigurationStores ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresListKeys.json
     */
    /**
     * Sample code: ConfigurationStores_ListKeys.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresListKeys(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().listKeys("myResourceGroup", "contoso", null, Context.NONE);
    }
}
