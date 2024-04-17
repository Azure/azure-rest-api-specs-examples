
/**
 * Samples for Routes ListByEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Routes_ListByEndpoint.json
     */
    /**
     * Sample code: Routes_ListByEndpoint.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routesListByEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getRoutes().listByEndpoint("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
