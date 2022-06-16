import com.azure.core.util.Context;

/** Samples for BlobServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/BlobServicesList.json
     */
    /**
     * Sample code: ListBlobServices.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listBlobServices(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobServices().list("res4410", "sto8607", Context.NONE);
    }
}
