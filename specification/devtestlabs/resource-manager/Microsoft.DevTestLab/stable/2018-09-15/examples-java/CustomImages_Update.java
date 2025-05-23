
import com.azure.resourcemanager.devtestlabs.models.CustomImage;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CustomImages Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/CustomImages_Update.
     * json
     */
    /**
     * Sample code: CustomImages_Update.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void customImagesUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        CustomImage resource = manager.customImages().getWithResponse("resourceGroupName", "{labName}",
            "{customImageName}", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("tagName1", "tagValue2")).apply();
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
