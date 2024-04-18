
/**
 * Samples for AfdEndpoints ListResourceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDEndpoints_ListResourceUsage.json
     */
    /**
     * Sample code: AFDEndpoints_ListResourceUsage.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDEndpointsListResourceUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdEndpoints().listResourceUsage("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
