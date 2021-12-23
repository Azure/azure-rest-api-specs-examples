Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-edgeorder_1.0.0-beta.1/sdk/edgeorder/azure-resourcemanager-edgeorder/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.edgeorder.models.FilterableProperty;
import com.azure.resourcemanager.edgeorder.models.ProductFamiliesRequest;
import com.azure.resourcemanager.edgeorder.models.SupportedFilterTypes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ResourceProvider ListProductFamilies. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListProductFamilies.json
     */
    /**
     * Sample code: ListProductFamilies.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listProductFamilies(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager
            .resourceProviders()
            .listProductFamilies(
                new ProductFamiliesRequest()
                    .withFilterableProperties(
                        mapOf(
                            "azurestackedge",
                            Arrays
                                .asList(
                                    new FilterableProperty()
                                        .withType(SupportedFilterTypes.SHIP_TO_COUNTRIES)
                                        .withSupportedValues(Arrays.asList("US"))))),
                null,
                null,
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
