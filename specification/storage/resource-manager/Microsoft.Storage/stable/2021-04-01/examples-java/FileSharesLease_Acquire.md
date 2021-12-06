Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.LeaseShareAction;
import com.azure.resourcemanager.storage.models.LeaseShareRequest;

/** Samples for FileShares Lease. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/FileSharesLease_Acquire.json
     */
    /**
     * Sample code: Acquire a lease on a share.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void acquireALeaseOnAShare(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .leaseWithResponse(
                "res3376",
                "sto328",
                "share124",
                null,
                new LeaseShareRequest().withAction(LeaseShareAction.ACQUIRE).withLeaseDuration(-1),
                Context.NONE);
    }
}
```
