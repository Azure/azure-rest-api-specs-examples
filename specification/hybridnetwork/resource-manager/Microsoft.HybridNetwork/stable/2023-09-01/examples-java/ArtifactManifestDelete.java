
/**
 * Samples for ArtifactManifests Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactManifestDelete.json
     */
    /**
     * Sample code: Delete a artifact manifest resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        deleteAArtifactManifestResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactManifests().delete("rg", "TestPublisher", "TestArtifactStore", "TestManifest",
            com.azure.core.util.Context.NONE);
    }
}
