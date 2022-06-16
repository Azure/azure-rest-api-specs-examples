import com.azure.core.util.Context;

/** Samples for BlobContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/BlobContainersGet.json
     */
    /**
     * Sample code: GetContainers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getContainers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getBlobContainers()
            .getWithResponse("res9871", "sto6217", "container1634", Context.NONE);
    }
}
