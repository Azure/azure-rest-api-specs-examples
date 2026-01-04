
/**
 * Samples for StaticCidrs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/StaticCidrs_Get.json
     */
    /**
     * Sample code: StaticCidrs_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void staticCidrsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getStaticCidrs().getWithResponse("rg1", "TestNetworkManager",
            "TestPool", "TestStaticCidr", com.azure.core.util.Context.NONE);
    }
}
