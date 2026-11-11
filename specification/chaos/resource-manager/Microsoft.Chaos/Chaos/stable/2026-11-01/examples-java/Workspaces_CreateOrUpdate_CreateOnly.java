
import com.azure.resourcemanager.chaos.models.WorkspaceProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workspaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Workspaces_CreateOrUpdate_CreateOnly.json
     */
    /**
     * Sample code: Create a workspace only when absent, or return 412 when it already exists.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createAWorkspaceOnlyWhenAbsentOrReturn412WhenItAlreadyExists(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().define("exampleWorkspace").withRegion("westus2").withExistingResourceGroup("exampleRG")
            .withProperties(new WorkspaceProperties().withScopes(
                Arrays.asList("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG")))
            .withIfNoneMatch("*").create();
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
