Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.ExtendedLocation;
import com.azure.resourcemanager.compute.models.ExtendedLocationTypes;

/** Samples for Disks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/CreateAnEmptyManagedDiskInExtendedLocation.json
     */
    /**
     * Sample code: Create an empty managed disk in extended location.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAnEmptyManagedDiskInExtendedLocation(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .createOrUpdate(
                "myResourceGroup",
                "myDisk",
                new DiskInner()
                    .withLocation("West US")
                    .withExtendedLocation(
                        new ExtendedLocation().withName("{edge-zone-id}").withType(ExtendedLocationTypes.EDGE_ZONE))
                    .withCreationData(new CreationData().withCreateOption(DiskCreateOption.EMPTY))
                    .withDiskSizeGB(200),
                Context.NONE);
    }
}
```
