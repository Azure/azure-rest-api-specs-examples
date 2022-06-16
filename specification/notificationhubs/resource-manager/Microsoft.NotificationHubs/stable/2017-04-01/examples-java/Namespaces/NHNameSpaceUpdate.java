import com.azure.core.util.Context;
import com.azure.resourcemanager.notificationhubs.models.NamespaceResource;
import com.azure.resourcemanager.notificationhubs.models.Sku;
import com.azure.resourcemanager.notificationhubs.models.SkuName;
import java.util.HashMap;
import java.util.Map;

/** Samples for Namespaces Patch. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceUpdate.json
     */
    /**
     * Sample code: NameSpaceUpdate.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceUpdate(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        NamespaceResource resource =
            manager.namespaces().getByResourceGroupWithResponse("5ktrial", "nh-sdk-ns", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withSku(new Sku().withName(SkuName.STANDARD).withTier("Standard"))
            .apply();
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
