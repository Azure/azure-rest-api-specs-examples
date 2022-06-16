import com.azure.core.util.Context;

/** Samples for ConfigurationStores Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresDelete.json
     */
    /**
     * Sample code: ConfigurationStores_Delete.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresDelete(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().delete("myResourceGroup", "contoso", Context.NONE);
    }
}
