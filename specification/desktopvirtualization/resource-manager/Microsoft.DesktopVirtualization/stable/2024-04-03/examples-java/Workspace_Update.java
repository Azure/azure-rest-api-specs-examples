
import com.azure.resourcemanager.desktopvirtualization.models.Workspace;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Workspaces Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * Workspace_Update.json
     */
    /**
     * Sample code: Workspace_Update.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void
        workspaceUpdate(com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        Workspace resource = manager.workspaces()
            .getByResourceGroupWithResponse("resourceGroup1", "workspace1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).withDescription("des1")
            .withFriendlyName("friendly").apply();
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
