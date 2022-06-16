import com.azure.core.util.Context;

/** Samples for ConfigurationStores List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresList.json
     */
    /**
     * Sample code: ConfigurationStores_List.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresList(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().list(null, Context.NONE);
    }
}
