
import com.azure.resourcemanager.storage.models.LeaseContainerRequest;
import com.azure.resourcemanager.storage.models.LeaseContainerRequestAction;

/**
 * Samples for BlobContainers Lease.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/BlobContainersLease_Break.json
     */
    /**
     * Sample code: Break a lease on a container.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void breakALeaseOnAContainer(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().leaseWithResponse(
            "res3376", "sto328", "container6185", new LeaseContainerRequest()
                .withAction(LeaseContainerRequestAction.BREAK).withLeaseId("8698f513-fa75-44a1-b8eb-30ba336af27d"),
            com.azure.core.util.Context.NONE);
    }
}
