
/**
 * Samples for Endpoints Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Endpoints_Stop.json
     */
    /**
     * Sample code: Endpoints_Stop.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsStop(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getEndpoints().stop("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
