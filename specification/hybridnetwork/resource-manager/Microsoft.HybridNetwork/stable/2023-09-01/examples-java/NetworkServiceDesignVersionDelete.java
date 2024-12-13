
/**
 * Samples for NetworkServiceDesignVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkServiceDesignVersionDelete.json
     */
    /**
     * Sample code: Delete a network service design version.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        deleteANetworkServiceDesignVersion(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkServiceDesignVersions().delete("rg", "TestPublisher", "TestNetworkServiceDesignGroupName",
            "1.0.0", com.azure.core.util.Context.NONE);
    }
}
