
/**
 * Samples for AvailablePrivateEndpointTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * AvailablePrivateEndpointTypesGet.json
     */
    /**
     * Sample code: Get available PrivateEndpoint types.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailablePrivateEndpointTypes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAvailablePrivateEndpointTypes().list("regionName",
            com.azure.core.util.Context.NONE);
    }
}
