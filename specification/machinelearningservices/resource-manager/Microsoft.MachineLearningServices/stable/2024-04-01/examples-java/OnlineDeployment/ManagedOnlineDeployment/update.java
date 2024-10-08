
import com.azure.resourcemanager.machinelearning.models.OnlineDeployment;
import com.azure.resourcemanager.machinelearning.models.PartialSku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for OnlineDeployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/OnlineDeployment/ManagedOnlineDeployment/update.json
     */
    /**
     * Sample code: Update Managed Online Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        updateManagedOnlineDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        OnlineDeployment resource = manager.onlineDeployments().getWithResponse("test-rg", "my-aml-workspace",
            "testEndpointName", "testDeploymentName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf()).withSku(new PartialSku().withName("string").withTier(SkuTier.FREE)
            .withSize("string").withFamily("string").withCapacity(1)).apply();
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
