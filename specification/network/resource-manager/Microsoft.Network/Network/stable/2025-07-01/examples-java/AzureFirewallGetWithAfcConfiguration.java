
/**
 * Samples for AzureFirewalls GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AzureFirewallGetWithAfcConfiguration.json
     */
    /**
     * Sample code: Get Azure Firewall With AFC Control Plane.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAzureFirewallWithAFCControlPlane(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAzureFirewalls().getByResourceGroupWithResponse("rg1", "azurefirewall",
            com.azure.core.util.Context.NONE);
    }
}
