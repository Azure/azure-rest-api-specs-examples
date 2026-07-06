
import com.azure.resourcemanager.storage.models.ListContainersInclude;

/**
 * Samples for BlobContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/DeletedBlobContainersList.json
     */
    /**
     * Sample code: ListDeletedContainers.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listDeletedContainers(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().list("res9290", "sto1590", null, null,
            ListContainersInclude.DELETED, com.azure.core.util.Context.NONE);
    }
}
