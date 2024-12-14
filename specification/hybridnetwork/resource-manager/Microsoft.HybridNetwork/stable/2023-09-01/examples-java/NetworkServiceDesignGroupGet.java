
/**
 * Samples for NetworkServiceDesignGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkServiceDesignGroupGet.json
     */
    /**
     * Sample code: Get a networkServiceDesign group resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        getANetworkServiceDesignGroupResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkServiceDesignGroups().getWithResponse("rg", "TestPublisher", "TestNetworkServiceDesignGroupName",
            com.azure.core.util.Context.NONE);
    }
}
