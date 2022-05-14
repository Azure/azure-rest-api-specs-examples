Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.PartnerConfigurationUpdateParameters;
import java.util.HashMap;
import java.util.Map;

/** Samples for PartnerConfigurations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerConfigurations_Update.json
     */
    /**
     * Sample code: PartnerConfigurations_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerConfigurationsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerConfigurations()
            .update(
                "examplerg",
                new PartnerConfigurationUpdateParameters()
                    .withTags(mapOf("tag1", "value11", "tag2", "value22"))
                    .withDefaultMaximumExpirationTimeInDays(100),
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
