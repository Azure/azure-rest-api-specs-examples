
import com.azure.resourcemanager.storagemover.models.Endpoint;

/**
 * Samples for Endpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Endpoints_Update_AzureStorageNfsFileShare.json
     */
    /**
     * Sample code: Endpoints_Update_AzureStorageNfsFileShare.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        endpointsUpdateAzureStorageNfsFileShare(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        Endpoint resource = manager.endpoints().getWithResponse("examples-rg", "examples-storageMoverName",
            "examples-endpointName", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
