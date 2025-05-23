
import com.azure.resourcemanager.devtestlabs.models.Policy;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Policies Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_Update.json
     */
    /**
     * Sample code: Policies_Update.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void policiesUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        Policy resource = manager.policies().getWithResponse("resourceGroupName", "{labName}", "{policySetName}",
            "{policyName}", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tagName1", "tagValue1")).apply();
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
