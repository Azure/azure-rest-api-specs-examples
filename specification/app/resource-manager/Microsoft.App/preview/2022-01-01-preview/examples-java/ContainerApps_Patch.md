Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appcontainers_1.0.0-beta.1/sdk/appcontainers/azure-resourcemanager-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appcontainers.models.ContainerApp;
import java.util.HashMap;
import java.util.Map;

/** Samples for ContainerApps Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/ContainerApps_Patch.json
     */
    /**
     * Sample code: Patch Container App.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void patchContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        ContainerApp resource =
            manager.containerApps().getByResourceGroupWithResponse("rg", "testcontainerApp0", Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
