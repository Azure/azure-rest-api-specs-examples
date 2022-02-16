Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.Resource;
import com.azure.core.util.Context;
import java.util.HashMap;
import java.util.Map;

/** Samples for ContainerGroups Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-09-01/examples/ContainerGroupsUpdate.json
     */
    /**
     * Sample code: ContainerGroupsUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerGroupsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerGroups()
            .manager()
            .serviceClient()
            .getContainerGroups()
            .updateWithResponse(
                "demoResource",
                "demo1",
                new Resource().withTags(mapOf("tag1key", "tag1Value", "tag2key", "tag2Value")),
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
