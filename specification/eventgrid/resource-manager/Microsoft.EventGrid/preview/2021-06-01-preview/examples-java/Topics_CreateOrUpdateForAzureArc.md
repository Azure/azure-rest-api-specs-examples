```java
import com.azure.resourcemanager.eventgrid.models.ExtendedLocation;
import com.azure.resourcemanager.eventgrid.models.InputSchema;
import com.azure.resourcemanager.eventgrid.models.ResourceKind;
import java.util.HashMap;
import java.util.Map;

/** Samples for Topics CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/Topics_CreateOrUpdateForAzureArc.json
     */
    /**
     * Sample code: Topics_CreateOrUpdateForAzureArc.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsCreateOrUpdateForAzureArc(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .topics()
            .define("exampletopic1")
            .withRegion("westus2")
            .withExistingResourceGroup("examplerg")
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withKind(ResourceKind.AZURE_ARC)
            .withExtendedLocation(
                new ExtendedLocation()
                    .withName(
                        "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.ExtendedLocation/CustomLocations/exampleCustomLocation")
                    .withType("CustomLocation"))
            .withInputSchema(InputSchema.CLOUD_EVENT_SCHEMA_V1_0)
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.1.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.
