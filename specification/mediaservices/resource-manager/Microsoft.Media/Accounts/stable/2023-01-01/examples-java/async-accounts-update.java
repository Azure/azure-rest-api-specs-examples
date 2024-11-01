
import com.azure.resourcemanager.mediaservices.models.MediaService;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Mediaservices Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/async-accounts-
     * update.json
     */
    /**
     * Sample code: Update a Media Services accounts.
     * 
     * @param manager Entry point to MediaServicesManager.
     */
    public static void
        updateAMediaServicesAccounts(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        MediaService resource = manager.mediaservices()
            .getByResourceGroupWithResponse("contosorg", "contososports", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder")).apply();
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
