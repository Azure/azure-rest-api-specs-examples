
/**
 * Samples for NetworkServiceDesignVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkServiceDesignVersionGet.json
     */
    /**
     * Sample code: Get a network service design version resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        getANetworkServiceDesignVersionResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkServiceDesignVersions().getWithResponse("rg", "TestPublisher",
            "TestNetworkServiceDesignGroupName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
