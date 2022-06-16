import com.azure.resourcemanager.notificationhubs.models.Sku;
import com.azure.resourcemanager.notificationhubs.models.SkuName;
import java.util.HashMap;
import java.util.Map;

/** Samples for Namespaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceCreate.json
     */
    /**
     * Sample code: NameSpaceCreate.
     *
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void nameSpaceCreate(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager
            .namespaces()
            .define("nh-sdk-ns")
            .withLocation("South Central US")
            .withExistingResourceGroup("5ktrial")
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withSku(new Sku().withName(SkuName.STANDARD).withTier("Standard"))
            .create();
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
