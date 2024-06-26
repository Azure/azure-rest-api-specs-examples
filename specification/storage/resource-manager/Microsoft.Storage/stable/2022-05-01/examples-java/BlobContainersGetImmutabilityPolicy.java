import com.azure.core.util.Context;

/** Samples for BlobContainers GetImmutabilityPolicy. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/BlobContainersGetImmutabilityPolicy.json
     */
    /**
     * Sample code: GetImmutabilityPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getImmutabilityPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .getImmutabilityPolicyWithResponse("res5221", "sto9177", "container3489", null, Context.NONE);
    }
}
