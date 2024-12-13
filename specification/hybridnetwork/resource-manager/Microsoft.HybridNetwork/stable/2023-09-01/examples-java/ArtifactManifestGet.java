
/**
 * Samples for ArtifactManifests Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactManifestGet.json
     */
    /**
     * Sample code: Get a artifact manifest resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        getAArtifactManifestResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactManifests().getWithResponse("rg", "TestPublisher", "TestArtifactStore", "TestManifest",
            com.azure.core.util.Context.NONE);
    }
}
