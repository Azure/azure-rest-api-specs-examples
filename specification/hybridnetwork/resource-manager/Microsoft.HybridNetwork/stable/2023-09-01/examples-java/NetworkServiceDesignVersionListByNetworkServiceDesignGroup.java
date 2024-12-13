
/**
 * Samples for NetworkServiceDesignVersions ListByNetworkServiceDesignGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkServiceDesignVersionListByNetworkServiceDesignGroup.json
     */
    /**
     * Sample code: Get Publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getPublisherResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkServiceDesignVersions().listByNetworkServiceDesignGroup("rg", "TestPublisher",
            "TestNetworkServiceDesignGroupName", com.azure.core.util.Context.NONE);
    }
}
