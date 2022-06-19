import com.azure.core.util.Context;

/** Samples for ConfigurationStores ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresListByResourceGroup.json
     */
    /**
     * Sample code: ConfigurationStores_ListByResourceGroup.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void configurationStoresListByResourceGroup(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.configurationStores().listByResourceGroup("myResourceGroup", null, Context.NONE);
    }
}
