
import com.azure.resourcemanager.storage.models.ListContainersInclude;

/** Samples for BlobContainers List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/DeletedBlobContainersList.
     * json
     */
    /**
     * Sample code: ListDeletedContainers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDeletedContainers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobContainers().list("res9290", "sto1590", null, null,
            ListContainersInclude.DELETED, com.azure.core.util.Context.NONE);
    }
}
