
import com.azure.resourcemanager.storage.models.LeaseContainerRequest;
import com.azure.resourcemanager.storage.models.LeaseContainerRequestAction;

/** Samples for BlobContainers Lease. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobContainersLease_Acquire.
     * json
     */
    /**
     * Sample code: Acquire a lease on a container.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void acquireALeaseOnAContainer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobContainers().leaseWithResponse("res3376", "sto328",
            "container6185",
            new LeaseContainerRequest().withAction(LeaseContainerRequestAction.ACQUIRE).withLeaseDuration(-1),
            com.azure.core.util.Context.NONE);
    }
}
