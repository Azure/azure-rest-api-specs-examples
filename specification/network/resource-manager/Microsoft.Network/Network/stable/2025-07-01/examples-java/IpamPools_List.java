
/**
 * Samples for IpamPools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpamPools_List.json
     */
    /**
     * Sample code: IpamPools_List.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void ipamPoolsList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().list("rg1", "TestNetworkManager", null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
