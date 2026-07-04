
/**
 * Samples for BastionHosts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BastionHostDeveloperGet.json
     */
    /**
     * Sample code: Get Developer Bastion Host.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getDeveloperBastionHost(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getBastionHosts().getByResourceGroupWithResponse("rg1", "bastionhostdeveloper'",
            com.azure.core.util.Context.NONE);
    }
}
