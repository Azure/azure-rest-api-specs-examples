
import com.azure.resourcemanager.storage.models.LeaseShareAction;
import com.azure.resourcemanager.storage.models.LeaseShareRequest;

/**
 * Samples for FileShares Lease.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileSharesLease_Acquire.json
     */
    /**
     * Sample code: Acquire a lease on a share.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void acquireALeaseOnAShare(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileShares().leaseWithResponse("res3376", "sto328", "share124", null,
            new LeaseShareRequest().withAction(LeaseShareAction.ACQUIRE).withLeaseDuration(-1),
            com.azure.core.util.Context.NONE);
    }
}
