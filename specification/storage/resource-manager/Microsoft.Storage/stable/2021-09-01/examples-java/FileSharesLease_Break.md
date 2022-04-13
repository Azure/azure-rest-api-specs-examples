Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.LeaseShareAction;
import com.azure.resourcemanager.storage.models.LeaseShareRequest;

/** Samples for FileShares Lease. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/FileSharesLease_Break.json
     */
    /**
     * Sample code: Break a lease on a share.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void breakALeaseOnAShare(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileShares()
            .leaseWithResponse(
                "res3376",
                "sto328",
                "share12",
                null,
                new LeaseShareRequest()
                    .withAction(LeaseShareAction.BREAK)
                    .withLeaseId("8698f513-fa75-44a1-b8eb-30ba336af27d"),
                Context.NONE);
    }
}
```
