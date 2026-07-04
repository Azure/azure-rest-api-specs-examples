
/**
 * Samples for IpamPools ListAssociatedResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpamPools_ListAssociatedResources.json
     */
    /**
     * Sample code: IpamPools_ListAssociatedResources.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void ipamPoolsListAssociatedResources(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().listAssociatedResources("rg1", "TestNetworkManager", "TestPool",
            com.azure.core.util.Context.NONE);
    }
}
