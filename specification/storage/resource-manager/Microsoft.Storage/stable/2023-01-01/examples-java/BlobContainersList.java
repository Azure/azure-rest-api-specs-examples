/** Samples for BlobContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/BlobContainersList.json
     */
    /**
     * Sample code: ListContainers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listContainers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .list("res9290", "sto1590", null, null, null, com.azure.core.util.Context.NONE);
    }
}
