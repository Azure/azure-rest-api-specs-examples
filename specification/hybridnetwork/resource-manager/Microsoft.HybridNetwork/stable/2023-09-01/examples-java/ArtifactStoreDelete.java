
/**
 * Samples for ArtifactStores Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactStoreDelete.json
     */
    /**
     * Sample code: Delete a artifact store of publisher resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        deleteAArtifactStoreOfPublisherResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactStores().delete("rg", "TestPublisher", "TestArtifactStore", com.azure.core.util.Context.NONE);
    }
}
