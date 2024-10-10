
/**
 * Samples for LinkersOperation ListDaprConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * GetDaprConfigurations.json
     */
    /**
     * Sample code: GetDaprConfigurations.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void getDaprConfigurations(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.linkersOperations().listDaprConfigurations(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
            com.azure.core.util.Context.NONE);
    }
}
