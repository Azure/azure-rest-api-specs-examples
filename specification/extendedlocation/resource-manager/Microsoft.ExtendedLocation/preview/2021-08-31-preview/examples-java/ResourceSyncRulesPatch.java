
import com.azure.resourcemanager.extendedlocation.models.ResourceSyncRule;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ResourceSyncRules Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/
     * ResourceSyncRulesPatch.json
     */
    /**
     * Sample code: Update Resource Sync Rule.
     * 
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void
        updateResourceSyncRule(com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        ResourceSyncRule resource = manager.resourceSyncRules().getWithResponse("testresourcegroup", "customLocation01",
            "resourceSyncRule01", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tier", "testing"))
            .withTargetResourceGroup("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/")
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
