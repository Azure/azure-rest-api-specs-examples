
/**
 * Samples for ArtifactStores ListByPublisher.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactStoresListByPublisherName.json
     */
    /**
     * Sample code: Get application groups under a publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getApplicationGroupsUnderAPublisherResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactStores().listByPublisher("rg", "TestPublisher", com.azure.core.util.Context.NONE);
    }
}
