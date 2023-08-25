/** Samples for BlobContainers LockImmutabilityPolicy. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobContainersLockImmutabilityPolicy.json
     */
    /**
     * Sample code: LockImmutabilityPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void lockImmutabilityPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .lockImmutabilityPolicyWithResponse(
                "res2702", "sto5009", "container1631", "8d59f825b721dd3", com.azure.core.util.Context.NONE);
    }
}
