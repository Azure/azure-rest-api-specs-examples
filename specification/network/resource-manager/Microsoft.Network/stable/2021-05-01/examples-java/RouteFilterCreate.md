Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.RouteFilterInner;
import com.azure.resourcemanager.network.fluent.models.RouteFilterRuleInner;
import com.azure.resourcemanager.network.models.Access;
import com.azure.resourcemanager.network.models.RouteFilterRuleType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for RouteFilters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RouteFilterCreate.json
     */
    /**
     * Sample code: RouteFilterCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routeFilterCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getRouteFilters()
            .createOrUpdate(
                "rg1",
                "filterName",
                new RouteFilterInner()
                    .withLocation("West US")
                    .withTags(mapOf("key1", "value1"))
                    .withRules(
                        Arrays
                            .asList(
                                new RouteFilterRuleInner()
                                    .withName("ruleName")
                                    .withAccess(Access.ALLOW)
                                    .withRouteFilterRuleType(RouteFilterRuleType.COMMUNITY)
                                    .withCommunities(Arrays.asList("12076:5030", "12076:5040")))),
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
