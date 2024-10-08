
import com.azure.resourcemanager.machinelearning.models.BuildContext;
import com.azure.resourcemanager.machinelearning.models.EnvironmentVersionProperties;
import com.azure.resourcemanager.machinelearning.models.InferenceContainerProperties;
import com.azure.resourcemanager.machinelearning.models.Route;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for EnvironmentVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/EnvironmentVersion/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Environment Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceEnvironmentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentVersions().define("string").withExistingEnvironment("test-rg", "my-aml-workspace", "string")
            .withProperties(new EnvironmentVersionProperties().withDescription("string")
                .withTags(mapOf("string", "string")).withProperties(mapOf("string", "string")).withIsAnonymous(false)
                .withImage("docker.io/tensorflow/serving:latest").withCondaFile("string")
                .withBuild(new BuildContext().withContextUri(
                    "https://storage-account.blob.core.windows.net/azureml/DockerBuildContext/95ddede6b9b8c4e90472db3acd0a8d28/")
                    .withDockerfilePath("prod/Dockerfile"))
                .withInferenceConfig(
                    new InferenceContainerProperties().withLivenessRoute(new Route().withPath("string").withPort(1))
                        .withReadinessRoute(new Route().withPath("string").withPort(1))
                        .withScoringRoute(new Route().withPath("string").withPort(1))))
            .create();
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
