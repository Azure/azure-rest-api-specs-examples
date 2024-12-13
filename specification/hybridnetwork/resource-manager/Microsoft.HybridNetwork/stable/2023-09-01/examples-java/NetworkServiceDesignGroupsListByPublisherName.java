
/**
 * Samples for NetworkServiceDesignGroups ListByPublisher.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkServiceDesignGroupsListByPublisherName.json
     */
    /**
     * Sample code: Get networkServiceDesign groups under publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getNetworkServiceDesignGroupsUnderPublisherResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkServiceDesignGroups().listByPublisher("rg", "TestPublisher", com.azure.core.util.Context.NONE);
    }
}
