
import com.azure.resourcemanager.hybridnetwork.models.ArtifactManifestPropertiesFormat;
import com.azure.resourcemanager.hybridnetwork.models.ArtifactType;
import com.azure.resourcemanager.hybridnetwork.models.ManifestArtifactFormat;
import java.util.Arrays;

/**
 * Samples for ArtifactManifests CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactManifestCreate.json
     */
    /**
     * Sample code: Create or update the artifact manifest resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateTheArtifactManifestResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.artifactManifests().define("TestManifest").withRegion("eastus")
            .withExistingArtifactStore("rg", "TestPublisher", "TestArtifactStore")
            .withProperties(new ArtifactManifestPropertiesFormat().withArtifacts(Arrays.asList(
                new ManifestArtifactFormat().withArtifactName("fed-rbac").withArtifactType(ArtifactType.OCIARTIFACT)
                    .withArtifactVersion("1.0.0"),
                new ManifestArtifactFormat().withArtifactName("nginx").withArtifactType(ArtifactType.OCIARTIFACT)
                    .withArtifactVersion("v1"))))
            .create();
    }
}
