
import com.azure.resourcemanager.machinelearning.models.Aks;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Compute CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Compute/createOrUpdate/BasicAKSCompute.json
     */
    /**
     * Sample code: Create an AKS Compute.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createAnAKSCompute(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.computes().define("compute123").withExistingWorkspace("testrg123", "workspaces123").withRegion("eastus")
            .withProperties(new Aks()).create();
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
