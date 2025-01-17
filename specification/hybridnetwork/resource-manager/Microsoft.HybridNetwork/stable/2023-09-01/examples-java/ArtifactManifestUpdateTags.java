
import com.azure.resourcemanager.hybridnetwork.models.ArtifactManifest;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ArtifactManifests Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ArtifactManifestUpdateTags.json
     */
    /**
     * Sample code: Update a artifact manifest resource tags.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        updateAArtifactManifestResourceTags(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        ArtifactManifest resource = manager.artifactManifests().getWithResponse("rg", "TestPublisher",
            "TestArtifactStore", "TestManifest", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
