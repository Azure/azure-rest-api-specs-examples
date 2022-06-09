```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.InboundIpRule;
import com.azure.resourcemanager.eventgrid.models.IpActionType;
import com.azure.resourcemanager.eventgrid.models.PublicNetworkAccess;
import com.azure.resourcemanager.eventgrid.models.Topic;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Topics Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/Topics_Update.json
     */
    /**
     * Sample code: Topics_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        Topic resource =
            manager.topics().getByResourceGroupWithResponse("examplerg", "exampletopic1", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .withInboundIpRules(
                Arrays
                    .asList(
                        new InboundIpRule().withIpMask("12.18.30.15").withAction(IpActionType.ALLOW),
                        new InboundIpRule().withIpMask("12.18.176.1").withAction(IpActionType.ALLOW)))
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.
