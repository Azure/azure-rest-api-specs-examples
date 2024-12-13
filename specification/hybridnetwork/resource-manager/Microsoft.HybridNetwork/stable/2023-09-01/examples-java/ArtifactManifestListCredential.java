
/**
 * Samples for ArtifactManifests ListCredential.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactManifestListCredential.json
     */
    /**
     * Sample code: List a credential for artifact manifest.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        listACredentialForArtifactManifest(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactManifests().listCredentialWithResponse("rg", "TestPublisher", "TestArtifactStore",
            "TestArtifactManifestName", com.azure.core.util.Context.NONE);
    }
}
