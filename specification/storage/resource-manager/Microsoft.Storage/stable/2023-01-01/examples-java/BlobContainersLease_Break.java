
import com.azure.resourcemanager.storage.models.LeaseContainerRequest;
import com.azure.resourcemanager.storage.models.LeaseContainerRequestAction;

/** Samples for BlobContainers Lease. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobContainersLease_Break.
     * json
     */
    /**
     * Sample code: Break a lease on a container.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void breakALeaseOnAContainer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobContainers().leaseWithResponse(
            "res3376", "sto328", "container6185", new LeaseContainerRequest()
                .withAction(LeaseContainerRequestAction.BREAK).withLeaseId("8698f513-fa75-44a1-b8eb-30ba336af27d"),
            com.azure.core.util.Context.NONE);
    }
}
