Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_1.1.0-beta.3/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.MediaService;
import java.util.HashMap;
import java.util.Map;

/** Samples for Mediaservices Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-update.json
     */
    /**
     * Sample code: Update a Media Services accounts.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateAMediaServicesAccounts(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        MediaService resource =
            manager.mediaservices().getByResourceGroupWithResponse("contoso", "contososports", Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "value3")).apply();
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
```
