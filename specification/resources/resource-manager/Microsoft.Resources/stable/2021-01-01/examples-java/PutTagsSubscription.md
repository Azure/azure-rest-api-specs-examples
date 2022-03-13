Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.resources.fluent.models.TagsResourceInner;
import com.azure.resourcemanager.resources.models.Tags;
import java.util.HashMap;
import java.util.Map;

/** Samples for TagOperations CreateOrUpdateAtScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PutTagsSubscription.json
     */
    /**
     * Sample code: Update tags on a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateTagsOnASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .serviceClient()
            .getTagOperations()
            .createOrUpdateAtScopeWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                new TagsResourceInner()
                    .withProperties(new Tags().withTags(mapOf("tagKey1", "tag-value-1", "tagKey2", "tag-value-2"))),
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
