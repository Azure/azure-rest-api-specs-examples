
import com.azure.resourcemanager.extendedlocation.models.MatchExpressionsProperties;
import com.azure.resourcemanager.extendedlocation.models.ResourceSyncRulePropertiesSelector;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ResourceSyncRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/
     * ResourceSyncRulesCreate_Update.json
     */
    /**
     * Sample code: Create/Update Resource Sync Rule.
     * 
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void
        createUpdateResourceSyncRule(com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.resourceSyncRules().define("resourceSyncRule01").withRegion("West US")
            .withExistingCustomLocation("testresourcegroup", "customLocation01").withPriority(999)
            .withSelector(new ResourceSyncRulePropertiesSelector()
                .withMatchExpressions(Arrays.asList(new MatchExpressionsProperties().withKey("fakeTokenPlaceholder")
                    .withOperator("In").withValues(Arrays.asList("value4"))))
                .withMatchLabels(mapOf("key1", "fakeTokenPlaceholder")))
            .withTargetResourceGroup(
                "/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup")
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
