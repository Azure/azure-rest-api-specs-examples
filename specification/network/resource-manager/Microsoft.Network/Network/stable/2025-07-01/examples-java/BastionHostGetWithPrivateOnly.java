
/**
 * Samples for BastionHosts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostGetWithPrivateOnly.json
     */
    /**
     * Sample code: Get Bastion Host With Private Only.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getBastionHostWithPrivateOnly(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().getByResourceGroupWithResponse("rg1", "bastionhosttenant",
            com.azure.core.util.Context.NONE);
    }
}
