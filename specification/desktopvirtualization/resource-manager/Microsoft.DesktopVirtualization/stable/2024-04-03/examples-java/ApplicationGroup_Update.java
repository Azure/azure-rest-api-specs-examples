
import com.azure.resourcemanager.desktopvirtualization.models.ApplicationGroup;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ApplicationGroups Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * ApplicationGroup_Update.json
     */
    /**
     * Sample code: ApplicationGroups_Update.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void
        applicationGroupsUpdate(com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        ApplicationGroup resource = manager.applicationGroups()
            .getByResourceGroupWithResponse("resourceGroup1", "applicationGroup1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).withDescription("des1")
            .withFriendlyName("friendly").withShowInFeed(true).apply();
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
