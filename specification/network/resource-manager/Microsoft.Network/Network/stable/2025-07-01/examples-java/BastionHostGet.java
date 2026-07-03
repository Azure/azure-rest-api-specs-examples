
/**
 * Samples for BastionHosts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostGet.json
     */
    /**
     * Sample code: Get Bastion Host.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getBastionHost(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().getByResourceGroupWithResponse("rg1", "bastionhosttenant'",
            com.azure.core.util.Context.NONE);
    }
}
