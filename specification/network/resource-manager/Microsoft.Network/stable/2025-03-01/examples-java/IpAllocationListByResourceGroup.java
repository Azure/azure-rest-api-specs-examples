
/**
 * Samples for IpAllocations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * IpAllocationListByResourceGroup.json
     */
    /**
     * Sample code: List IpAllocations in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listIpAllocationsInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpAllocations().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
