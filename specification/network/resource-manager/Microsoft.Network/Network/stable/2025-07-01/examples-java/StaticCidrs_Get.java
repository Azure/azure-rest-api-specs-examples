
/**
 * Samples for StaticCidrs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/StaticCidrs_Get.json
     */
    /**
     * Sample code: StaticCidrs_Get.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void staticCidrsGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getStaticCidrs().getWithResponse("rg1", "TestNetworkManager", "TestPool",
            "TestStaticCidr", com.azure.core.util.Context.NONE);
    }
}
