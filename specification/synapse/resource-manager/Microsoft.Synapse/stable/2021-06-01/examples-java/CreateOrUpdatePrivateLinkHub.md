Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import java.util.HashMap;
import java.util.Map;

/** Samples for PrivateLinkHubs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdatePrivateLinkHub.json
     */
    /**
     * Sample code: Create or update a privateLinkHub.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateAPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .privateLinkHubs()
            .define("privateLinkHub1")
            .withRegion("East US")
            .withExistingResourceGroup("resourceGroup1")
            .withTags(mapOf("key", "value"))
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
