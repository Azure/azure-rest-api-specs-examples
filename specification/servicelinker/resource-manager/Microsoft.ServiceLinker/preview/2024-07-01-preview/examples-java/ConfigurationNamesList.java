
/**
 * Samples for ConfigurationNamesOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * ConfigurationNamesList.json
     */
    /**
     * Sample code: GetConfigurationNames.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void getConfigurationNames(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.configurationNamesOperations().list(null, null, com.azure.core.util.Context.NONE);
    }
}
