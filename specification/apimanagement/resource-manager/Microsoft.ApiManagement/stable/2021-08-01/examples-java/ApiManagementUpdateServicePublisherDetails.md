Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.ApiManagementServiceResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApiManagementService Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateServicePublisherDetails.json
     */
    /**
     * Sample code: ApiManagementUpdateServicePublisherDetails.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateServicePublisherDetails(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiManagementServiceResource resource =
            manager
                .apiManagementServices()
                .getByResourceGroupWithResponse("rg1", "apimService1", Context.NONE)
                .getValue();
        resource.update().withPublisherEmail("foobar@live.com").withPublisherName("Contoso Vnext").apply();
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
