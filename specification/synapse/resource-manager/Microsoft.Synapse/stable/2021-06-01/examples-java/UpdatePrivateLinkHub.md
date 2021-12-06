Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.PrivateLinkHub;
import java.util.HashMap;
import java.util.Map;

/** Samples for PrivateLinkHubs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdatePrivateLinkHub.json
     */
    /**
     * Sample code: Update a privateLinkHub.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateAPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        PrivateLinkHub resource =
            manager
                .privateLinkHubs()
                .getByResourceGroupWithResponse("resourceGroup1", "privateLinkHub1", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("key", "value")).apply();
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
