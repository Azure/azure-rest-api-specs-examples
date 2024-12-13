
/**
 * Samples for ArtifactManifests ListByArtifactStore.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactManifestListByArtifactStore.json
     */
    /**
     * Sample code: Get artifact manifest list resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        getArtifactManifestListResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactManifests().listByArtifactStore("rg", "TestPublisher", "TestArtifactStore",
            com.azure.core.util.Context.NONE);
    }
}
