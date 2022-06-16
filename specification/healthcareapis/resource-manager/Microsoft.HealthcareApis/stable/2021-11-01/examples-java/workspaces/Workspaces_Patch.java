import com.azure.core.util.Context;
import com.azure.resourcemanager.healthcareapis.models.Workspace;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/workspaces/Workspaces_Patch.json
     */
    /**
     * Sample code: Update a workspace.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void updateAWorkspace(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        Workspace resource =
            manager.workspaces().getByResourceGroupWithResponse("testRG", "workspace1", Context.NONE).getValue();
        resource.update().withTags(mapOf("tagKey", "tagValue")).apply();
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
