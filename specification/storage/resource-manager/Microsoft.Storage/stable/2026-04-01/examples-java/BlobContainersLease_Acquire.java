
import com.azure.resourcemanager.storage.models.LeaseContainerRequest;
import com.azure.resourcemanager.storage.models.LeaseContainerRequestAction;

/**
 * Samples for BlobContainers Lease.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobContainersLease_Acquire.json
     */
    /**
     * Sample code: Acquire a lease on a container.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void acquireALeaseOnAContainer(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().leaseWithResponse("res3376", "sto328", "container6185",
            new LeaseContainerRequest().withAction(LeaseContainerRequestAction.ACQUIRE).withLeaseDuration(-1),
            com.azure.core.util.Context.NONE);
    }
}
