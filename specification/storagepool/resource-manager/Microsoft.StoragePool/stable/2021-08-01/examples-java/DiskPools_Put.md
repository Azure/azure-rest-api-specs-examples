```java
import com.azure.resourcemanager.storagepool.fluent.models.Sku;
import com.azure.resourcemanager.storagepool.models.Disk;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for DiskPools CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Put.json
     */
    /**
     * Sample code: Create or Update Disk pool.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void createOrUpdateDiskPool(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager
            .diskPools()
            .define("myDiskPool")
            .withRegion("westus")
            .withExistingResourceGroup("myResourceGroup")
            .withSku(new Sku().withName("Basic_V1").withTier("Basic"))
            .withSubnetId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet")
            .withTags(mapOf("key", "value"))
            .withAvailabilityZones(Arrays.asList("1"))
            .withDisks(
                Arrays
                    .asList(
                        new Disk()
                            .withId(
                                "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_0"),
                        new Disk()
                            .withId(
                                "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1")))
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagepool_1.0.0-beta.1/sdk/storagepool/azure-resourcemanager-storagepool/README.md) on how to add the SDK to your project and authenticate.
