
/**
 * Samples for NetworkServiceDesignGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkServiceDesignGroupCreate.json
     */
    /**
     * Sample code: Create or update the network service design group.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateTheNetworkServiceDesignGroup(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkServiceDesignGroups().define("TestNetworkServiceDesignGroupName").withRegion("eastus")
            .withExistingPublisher("rg", "TestPublisher").create();
    }
}
