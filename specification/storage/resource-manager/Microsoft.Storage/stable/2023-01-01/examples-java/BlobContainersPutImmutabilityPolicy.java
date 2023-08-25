import com.azure.resourcemanager.storage.fluent.models.ImmutabilityPolicyInner;

/** Samples for BlobContainers CreateOrUpdateImmutabilityPolicy. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobContainersPutImmutabilityPolicy.json
     */
    /**
     * Sample code: CreateOrUpdateImmutabilityPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateImmutabilityPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .createOrUpdateImmutabilityPolicyWithResponse(
                "res1782",
                "sto7069",
                "container6397",
                null,
                new ImmutabilityPolicyInner()
                    .withImmutabilityPeriodSinceCreationInDays(3)
                    .withAllowProtectedAppendWrites(true),
                com.azure.core.util.Context.NONE);
    }
}
