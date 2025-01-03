
/**
 * Samples for Linker ListConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * GetConfigurations.json
     */
    /**
     * Sample code: GetConfiguration.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void getConfiguration(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.linkers().listConfigurationsWithResponse(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.App/containerApps/test-app",
            "linkName", com.azure.core.util.Context.NONE);
    }
}
