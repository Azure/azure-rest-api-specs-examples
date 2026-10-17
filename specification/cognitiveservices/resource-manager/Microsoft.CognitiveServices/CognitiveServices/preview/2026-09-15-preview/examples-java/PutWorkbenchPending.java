
import com.azure.resourcemanager.cognitiveservices.models.WorkbenchProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workbenches CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/PutWorkbenchPending.json
     */
    /**
     * Sample code: PutWorkbenchPending.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        putWorkbenchPending(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.workbenches().define("myWorkbench").withExistingProject("rgcognitiveservices", "myAccount", "myProject")
            .withProperties(new WorkbenchProperties().withTargetClusterId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.CognitiveServices/accounts/myAccount/computes/myCluster")
                .withImageLink("mcr.microsoft.com/azureml/curated/pytorch-gpu:latest"))
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
