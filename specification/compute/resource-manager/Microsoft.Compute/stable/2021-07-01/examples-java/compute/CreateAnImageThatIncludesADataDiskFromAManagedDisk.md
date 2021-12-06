Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.ImageInner;
import com.azure.resourcemanager.compute.models.ImageDataDisk;
import com.azure.resourcemanager.compute.models.ImageOSDisk;
import com.azure.resourcemanager.compute.models.ImageStorageProfile;
import com.azure.resourcemanager.compute.models.OperatingSystemStateTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;
import java.util.Arrays;

/** Samples for Images CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/CreateAnImageThatIncludesADataDiskFromAManagedDisk.json
     */
    /**
     * Sample code: Create a virtual machine image that includes a data disk from a managed disk.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVirtualMachineImageThatIncludesADataDiskFromAManagedDisk(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getImages()
            .createOrUpdate(
                "myResourceGroup",
                "myImage",
                new ImageInner()
                    .withLocation("West US")
                    .withStorageProfile(
                        new ImageStorageProfile()
                            .withOsDisk(
                                new ImageOSDisk()
                                    .withManagedDisk(
                                        new SubResource()
                                            .withId(
                                                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"))
                                    .withOsType(OperatingSystemTypes.LINUX)
                                    .withOsState(OperatingSystemStateTypes.GENERALIZED))
                            .withDataDisks(
                                Arrays
                                    .asList(
                                        new ImageDataDisk()
                                            .withManagedDisk(
                                                new SubResource()
                                                    .withId(
                                                        "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2"))
                                            .withLun(1)))
                            .withZoneResilient(false)),
                Context.NONE);
    }
}
```
