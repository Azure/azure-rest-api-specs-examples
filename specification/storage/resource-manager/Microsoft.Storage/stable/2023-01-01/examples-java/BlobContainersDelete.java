
/** Samples for BlobContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobContainersDelete.json
     */
    /**
     * Sample code: DeleteContainers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteContainers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobContainers().deleteWithResponse("res4079", "sto4506",
            "container9689", com.azure.core.util.Context.NONE);
    }
}
