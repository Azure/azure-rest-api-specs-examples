
import com.azure.resourcemanager.hybridnetwork.fluent.models.ArtifactManifestUpdateStateInner;
import com.azure.resourcemanager.hybridnetwork.models.ArtifactManifestState;

/**
 * Samples for ArtifactManifests UpdateState.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactManifestUpdateState.json
     */
    /**
     * Sample code: Update artifact manifest state.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        updateArtifactManifestState(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactManifests().updateState("rg", "TestPublisher", "TestArtifactStore", "TestArtifactManifestName",
            new ArtifactManifestUpdateStateInner().withArtifactManifestState(ArtifactManifestState.UPLOADED),
            com.azure.core.util.Context.NONE);
    }
}
