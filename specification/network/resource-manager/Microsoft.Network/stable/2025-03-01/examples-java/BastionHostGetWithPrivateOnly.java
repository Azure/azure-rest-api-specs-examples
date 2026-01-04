
/**
 * Samples for BastionHosts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/BastionHostGetWithPrivateOnly
     * .json
     */
    /**
     * Sample code: Get Bastion Host With Private Only.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getBastionHostWithPrivateOnly(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getBastionHosts().getByResourceGroupWithResponse("rg1",
            "bastionhosttenant", com.azure.core.util.Context.NONE);
    }
}
