
/**
 * Samples for PrivateEndpoints GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PrivateEndpointGet.json
     */
    /**
     * Sample code: Get private endpoint.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateEndpoints().getByResourceGroupWithResponse("rg1", "testPe",
            null, com.azure.core.util.Context.NONE);
    }
}
