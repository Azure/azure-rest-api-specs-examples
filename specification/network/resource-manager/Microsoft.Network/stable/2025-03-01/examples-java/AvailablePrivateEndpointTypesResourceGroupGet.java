
/**
 * Samples for AvailablePrivateEndpointTypes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * AvailablePrivateEndpointTypesResourceGroupGet.json
     */
    /**
     * Sample code: Get available PrivateEndpoint types in the resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAvailablePrivateEndpointTypesInTheResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAvailablePrivateEndpointTypes().listByResourceGroup("regionName",
            "rg1", com.azure.core.util.Context.NONE);
    }
}
