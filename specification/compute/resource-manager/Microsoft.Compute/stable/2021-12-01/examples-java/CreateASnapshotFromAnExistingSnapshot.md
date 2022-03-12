Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.SnapshotInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/** Samples for Snapshots CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/CreateASnapshotFromAnExistingSnapshot.json
     */
    /**
     * Sample code: Create a snapshot from an existing snapshot in the same or a different subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createASnapshotFromAnExistingSnapshotInTheSameOrADifferentSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .createOrUpdate(
                "myResourceGroup",
                "mySnapshot2",
                new SnapshotInner()
                    .withLocation("West US")
                    .withCreationData(
                        new CreationData()
                            .withCreateOption(DiskCreateOption.COPY)
                            .withSourceResourceId(
                                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1")),
                Context.NONE);
    }
}
```
