Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mediaservices_2.0.0/sdk/mediaservices/azure-resourcemanager-mediaservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.mediaservices.models.StreamingEndpoint;
import java.util.HashMap;
import java.util.Map;

/** Samples for StreamingEndpoints Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/streamingendpoint-update.json
     */
    /**
     * Sample code: Update a streaming endpoint.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void updateAStreamingEndpoint(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        StreamingEndpoint resource =
            manager
                .streamingEndpoints()
                .getWithResponse("mediaresources", "slitestmedia10", "myStreamingEndpoint1", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("tag3", "value3", "tag5", "value5"))
            .withDescription("test event 2")
            .withScaleUnits(5)
            .withAvailabilitySetName("availableset")
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
```
