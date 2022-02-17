Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/** Samples for Disks GrantAccess. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/BeginGetAccessManagedDiskWithVMGuestState.json
     */
    /**
     * Sample code: Get sas on managed disk and VM guest state.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSasOnManagedDiskAndVMGuestState(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .grantAccess(
                "myResourceGroup",
                "myDisk",
                new GrantAccessData()
                    .withAccess(AccessLevel.READ)
                    .withDurationInSeconds(300)
                    .withGetSecureVMGuestStateSas(true),
                Context.NONE);
    }
}
```
