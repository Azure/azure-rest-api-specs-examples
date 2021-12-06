Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagepool_1.0.0-beta.1/sdk/storagepool/azure-resourcemanager-storagepool/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storagepool.fluent.models.Sku;
import com.azure.resourcemanager.storagepool.models.Disk;
import com.azure.resourcemanager.storagepool.models.DiskPool;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for DiskPools Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Patch.json
     */
    /**
     * Sample code: Update Disk pool.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void updateDiskPool(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        DiskPool resource =
            manager
                .diskPools()
                .getByResourceGroupWithResponse("myResourceGroup", "myDiskPool", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf("key", "value"))
            .withSku(new Sku().withName("Basic_B1").withTier("Basic"))
            .withDisks(
                Arrays
                    .asList(
                        new Disk()
                            .withId(
                                "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_0"),
                        new Disk()
                            .withId(
                                "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1")))
            .apply();
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
