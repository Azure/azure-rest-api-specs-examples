Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.SnapshotInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/** Samples for Snapshots CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/CreateASnapshotByImportingAnUnmanagedBlobFromADifferentSubscription.json
     */
    /**
     * Sample code: Create a snapshot by importing an unmanaged blob from a different subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createASnapshotByImportingAnUnmanagedBlobFromADifferentSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .createOrUpdate(
                "myResourceGroup",
                "mySnapshot1",
                new SnapshotInner()
                    .withLocation("West US")
                    .withCreationData(
                        new CreationData()
                            .withCreateOption(DiskCreateOption.IMPORT)
                            .withStorageAccountId(
                                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount")
                            .withSourceUri("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd")),
                Context.NONE);
    }
}
```
