
import com.azure.resourcemanager.extendedlocation.models.CustomLocationFindTargetResourceGroupProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CustomLocations FindTargetResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/
     * CustomLocationsFindTargetResourceGroup.json
     */
    /**
     * Sample code: Post Custom Location Find Target Resource Group.
     * 
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void postCustomLocationFindTargetResourceGroup(
        com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.customLocations().findTargetResourceGroupWithResponse("testresourcegroup", "customLocation01",
            new CustomLocationFindTargetResourceGroupProperties()
                .withLabels(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
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
