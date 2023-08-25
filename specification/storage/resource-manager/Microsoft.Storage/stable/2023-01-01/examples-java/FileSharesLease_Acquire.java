import com.azure.resourcemanager.storage.models.LeaseShareAction;
import com.azure.resourcemanager.storage.models.LeaseShareRequest;

/** Samples for FileShares Lease. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/FileSharesLease_Acquire.json
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
                com.azure.core.util.Context.NONE);
    }
}
