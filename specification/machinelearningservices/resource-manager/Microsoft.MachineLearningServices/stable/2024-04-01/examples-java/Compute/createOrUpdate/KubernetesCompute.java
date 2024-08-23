
import com.azure.resourcemanager.machinelearning.models.InstanceTypeSchema;
import com.azure.resourcemanager.machinelearning.models.InstanceTypeSchemaResources;
import com.azure.resourcemanager.machinelearning.models.Kubernetes;
import com.azure.resourcemanager.machinelearning.models.KubernetesProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Compute CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/createOrUpdate/KubernetesCompute.json
     */
    /**
     * Sample code: Attach a Kubernetes Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        attachAKubernetesCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().define("compute123").withExistingWorkspace("testrg123", "workspaces123").withRegion("eastus")
            .withProperties(new Kubernetes().withDescription("some compute").withResourceId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/testrg123/providers/Microsoft.ContainerService/managedClusters/compute123-56826-c9b00420020b2")
                .withProperties(
                    new KubernetesProperties().withNamespace("default").withDefaultInstanceType("defaultInstanceType")
                        .withInstanceTypes(mapOf("defaultInstanceType",
                            new InstanceTypeSchema().withResources(new InstanceTypeSchemaResources()
                                .withRequests(mapOf("cpu", "1", "memory", "4Gi", "nvidia.com/gpu", null))
                                .withLimits(mapOf("cpu", "1", "memory", "4Gi", "nvidia.com/gpu", null)))))))
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
