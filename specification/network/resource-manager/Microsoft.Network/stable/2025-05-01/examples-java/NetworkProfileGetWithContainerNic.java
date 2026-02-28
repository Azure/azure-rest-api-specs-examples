
/**
 * Samples for NetworkProfiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkProfileGetWithContainerNic.json
     */
    /**
     * Sample code: Get network profile with container network interfaces.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getNetworkProfileWithContainerNetworkInterfaces(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkProfiles().getByResourceGroupWithResponse("rg1",
            "networkProfile1", null, com.azure.core.util.Context.NONE);
    }
}
