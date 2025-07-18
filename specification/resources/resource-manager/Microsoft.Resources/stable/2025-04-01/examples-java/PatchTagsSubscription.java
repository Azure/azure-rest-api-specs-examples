
import com.azure.resourcemanager.resources.models.Tags;
import com.azure.resourcemanager.resources.models.TagsPatchOperation;
import com.azure.resourcemanager.resources.models.TagsPatchResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for TagOperations UpdateAtScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2025-04-01/examples/PatchTagsSubscription.
     * json
     */
    /**
     * Sample code: Update tags on a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTagsOnASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getTagOperations().updateAtScope(
            "subscriptions/00000000-0000-0000-0000-000000000000",
            new TagsPatchResource().withOperation(TagsPatchOperation.REPLACE).withProperties(
                new Tags().withTags(mapOf("tagKey1", "fakeTokenPlaceholder", "tagKey2", "fakeTokenPlaceholder"))),
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
