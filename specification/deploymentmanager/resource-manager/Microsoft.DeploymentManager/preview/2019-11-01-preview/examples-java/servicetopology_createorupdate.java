import java.util.HashMap;
import java.util.Map;

/** Samples for ServiceTopologies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopology_createorupdate.json
     */
    /**
     * Sample code: Create a topology with Artifact Source.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void createATopologyWithArtifactSource(
        com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .serviceTopologies()
            .define("myTopology")
            .withRegion("centralus")
            .withExistingResourceGroup("myResourceGroup")
            .withTags(mapOf())
            .withArtifactSourceId("Microsoft.DeploymentManager/artifactSources/myArtifactSource")
            .create();
    }

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
