
/**
 * Samples for NetworkSecurityPerimeterConfigurations ListByConfigurationStore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/ConfigurationStoresListNetworkSecurityPerimeterConfigurations.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurations_List.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void networkSecurityPerimeterConfigurationsList(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.networkSecurityPerimeterConfigurations().listByConfigurationStore("myResourceGroup", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
