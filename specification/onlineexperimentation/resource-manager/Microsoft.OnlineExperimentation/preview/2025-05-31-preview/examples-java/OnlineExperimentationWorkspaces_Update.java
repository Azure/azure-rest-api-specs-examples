
import com.azure.resourcemanager.onlineexperimentation.models.ManagedServiceIdentity;
import com.azure.resourcemanager.onlineexperimentation.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.onlineexperimentation.models.OnlineExperimentationWorkspace;
import com.azure.resourcemanager.onlineexperimentation.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for OnlineExperimentationWorkspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_Update.json
     */
    /**
     * Sample code: Update an Online Experimentation Workspace.
     * 
     * @param manager Entry point to OnlineExperimentationManager.
     */
    public static void updateAnOnlineExperimentationWorkspace(
        com.azure.resourcemanager.onlineexperimentation.OnlineExperimentationManager manager) {
        OnlineExperimentationWorkspace resource = manager.onlineExperimentationWorkspaces()
            .getByResourceGroupWithResponse("res9871", "expworkspace3", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("newKey", "fakeTokenPlaceholder")).withIdentity(new ManagedServiceIdentity()
            .withType(ManagedServiceIdentityType.USER_ASSIGNED)
            .withUserAssignedIdentities(mapOf(
                "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                new UserAssignedIdentity(),
                "/subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2",
                new UserAssignedIdentity())))
            .apply();
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
