```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.models.LeaseContainerRequest;
import com.azure.resourcemanager.storage.models.LeaseContainerRequestAction;

/** Samples for BlobContainers Lease. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersLease_Break.json
     */
    /**
     * Sample code: Break a lease on a container.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void breakALeaseOnAContainer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .leaseWithResponse(
                "res3376",
                "sto328",
                "container6185",
                new LeaseContainerRequest()
                    .withAction(LeaseContainerRequestAction.BREAK)
                    .withLeaseId("8698f513-fa75-44a1-b8eb-30ba336af27d"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
