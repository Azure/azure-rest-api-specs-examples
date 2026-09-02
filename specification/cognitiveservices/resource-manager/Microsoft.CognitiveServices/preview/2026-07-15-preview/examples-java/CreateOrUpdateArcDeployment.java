
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentComputeType;
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentCpuMemoryResourceRequirements;
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentKubernetesResources;
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentModel;
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentProperties;
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentResourceRequirements;
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentRuntime;
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentSku;
import com.azure.resourcemanager.cognitiveservices.models.ArcDeploymentSkuName;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ArcDeployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15-preview/CreateOrUpdateArcDeployment.json
     */
    /**
     * Sample code: CreateOrUpdateArcDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        createOrUpdateArcDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.arcDeployments().define("phi-3-arc").withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new ArcDeploymentProperties()
                .withModel(new ArcDeploymentModel().withFormat("OpenAI").withName("phi-3-mini"))
                .withExtensionId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/edge-cluster/providers/Microsoft.KubernetesConfiguration/extensions/inference-operator")
                .withRuntime(ArcDeploymentRuntime.ONNX).withCompute(ArcDeploymentComputeType.CPU).withReplicas(2)
                .withResources(new ArcDeploymentKubernetesResources()
                    .withRequests(new ArcDeploymentCpuMemoryResourceRequirements().withCpu("8").withMemory("16Gi"))
                    .withLimits(new ArcDeploymentResourceRequirements().withCpu("8").withMemory("16Gi")))
                .withNodeSelector(mapOf("agentpool", "cpu")))
            .withSku(new ArcDeploymentSku().withName(ArcDeploymentSkuName.ARC)).create();
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
