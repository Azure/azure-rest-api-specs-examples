import com.azure.core.util.Context;

/** Samples for BlobContainers DeleteImmutabilityPolicy. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/BlobContainersDeleteImmutabilityPolicy.json
     */
    /**
     * Sample code: DeleteImmutabilityPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteImmutabilityPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .deleteImmutabilityPolicyWithResponse(
                "res1581", "sto9621", "container4910", "\"8d59f81a7fa7be0\"", Context.NONE);
    }
}
