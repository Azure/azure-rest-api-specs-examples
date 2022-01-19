Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.AccessLevel;
import com.azure.resourcemanager.compute.models.GrantAccessData;

/** Samples for Snapshots GrantAccess. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/BeginGetAccessSnapshot.json
     */
    /**
     * Sample code: Get a sas on a snapshot.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASasOnASnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getSnapshots()
            .grantAccess(
                "myResourceGroup",
                "mySnapshot",
                new GrantAccessData().withAccess(AccessLevel.READ).withDurationInSeconds(300),
                Context.NONE);
    }
}
```
