Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DiskRestorePoint RevokeAccess. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-12-01/examples/EndGetAccessDiskRestorePoint.json
     */
    /**
     * Sample code: Revokes access to a diskRestorePoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void revokesAccessToADiskRestorePoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskRestorePoints()
            .revokeAccess(
                "myResourceGroup",
                "rpc",
                "vmrp",
                "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745",
                Context.NONE);
    }
}
```
