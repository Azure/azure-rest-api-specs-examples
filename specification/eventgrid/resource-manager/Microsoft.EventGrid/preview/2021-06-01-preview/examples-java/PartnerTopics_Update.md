Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.PartnerTopicUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/** Samples for PartnerTopics Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/PartnerTopics_Update.json
     */
    /**
     * Sample code: PartnerTopics_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerTopics()
            .updateWithResponse(
                "examplerg",
                "examplePartnerTopicName1",
                new PartnerTopicUpdateParameters().withTags(mapOf("tag1", "value1", "tag2", "value2")),
                Context.NONE);
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
