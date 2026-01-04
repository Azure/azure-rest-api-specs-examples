
/**
 * Samples for IpamPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/IpamPools_Get.json
     */
    /**
     * Sample code: IpamPools_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void ipamPoolsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getIpamPools().getWithResponse("rg1", "TestNetworkManager",
            "TestPool", com.azure.core.util.Context.NONE);
    }
}
