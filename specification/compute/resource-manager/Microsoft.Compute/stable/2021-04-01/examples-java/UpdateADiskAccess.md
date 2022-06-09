```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.DiskAccessUpdate;
import java.util.HashMap;
import java.util.Map;

/** Samples for DiskAccesses Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/UpdateADiskAccess.json
     */
    /**
     * Sample code: Update a disk access resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateADiskAccessResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .update(
                "myResourceGroup",
                "myDiskAccess",
                new DiskAccessUpdate().withTags(mapOf("department", "Development", "project", "PrivateEndpoints")),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
