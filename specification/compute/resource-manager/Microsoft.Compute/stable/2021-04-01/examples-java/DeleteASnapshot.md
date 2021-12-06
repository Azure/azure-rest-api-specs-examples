Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Snapshots Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-04-01/examples/DeleteASnapshot.json
     */
    /**
     * Sample code: Delete a snapshot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .delete("myResourceGroup", "mySnapshot", Context.NONE);
    }
}
```
