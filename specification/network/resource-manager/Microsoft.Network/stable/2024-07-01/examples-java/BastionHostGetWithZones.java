
/**
 * Samples for BastionHosts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/BastionHostGetWithZones.json
     */
    /**
     * Sample code: Get Bastion Host With Zones.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getBastionHostWithZones(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getBastionHosts().getByResourceGroupWithResponse("rg1",
            "bastionhosttenant", com.azure.core.util.Context.NONE);
    }
}
