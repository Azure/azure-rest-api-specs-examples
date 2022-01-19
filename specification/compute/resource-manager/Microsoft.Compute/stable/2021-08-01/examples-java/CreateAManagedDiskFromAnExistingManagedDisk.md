Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/** Samples for Disks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/CreateAManagedDiskFromAnExistingManagedDisk.json
     */
    /**
     * Sample code: Create a managed disk from an existing managed disk in the same or different subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAManagedDiskFromAnExistingManagedDiskInTheSameOrDifferentSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .createOrUpdate(
                "myResourceGroup",
                "myDisk2",
                new DiskInner()
                    .withLocation("West US")
                    .withCreationData(
                        new CreationData()
                            .withCreateOption(DiskCreateOption.COPY)
                            .withSourceResourceId(
                                "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk1")),
                Context.NONE);
    }
}
```
