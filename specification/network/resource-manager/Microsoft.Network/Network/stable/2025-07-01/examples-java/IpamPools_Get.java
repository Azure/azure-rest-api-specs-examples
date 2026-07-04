
/**
 * Samples for IpamPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpamPools_Get.json
     */
    /**
     * Sample code: IpamPools_Get.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void ipamPoolsGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().getWithResponse("rg1", "TestNetworkManager", "TestPool",
            com.azure.core.util.Context.NONE);
    }
}
