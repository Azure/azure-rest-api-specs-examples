
import com.azure.resourcemanager.storage.models.LeaseShareAction;
import com.azure.resourcemanager.storage.models.LeaseShareRequest;

/** Samples for FileShares Lease. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/FileSharesLease_Break.json
     */
    /**
     * Sample code: Break a lease on a share.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void breakALeaseOnAShare(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileShares()
            .leaseWithResponse("res3376", "sto328", "share12", null, new LeaseShareRequest()
                .withAction(LeaseShareAction.BREAK).withLeaseId("8698f513-fa75-44a1-b8eb-30ba336af27d"),
                com.azure.core.util.Context.NONE);
    }
}
