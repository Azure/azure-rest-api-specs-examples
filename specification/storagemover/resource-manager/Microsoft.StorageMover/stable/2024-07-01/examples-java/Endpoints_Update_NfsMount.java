
import com.azure.resourcemanager.storagemover.models.Endpoint;

/**
 * Samples for Endpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/
     * Endpoints_Update_NfsMount.json
     */
    /**
     * Sample code: Endpoints_Update_NfsMount.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsUpdateNfsMount(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        Endpoint resource = manager.endpoints().getWithResponse("examples-rg", "examples-storageMoverName",
            "examples-endpointName", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }

    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/
     * Endpoints_Update_AzureStorageSmbFileShare.json
     */
    /**
     * Sample code: Endpoints_Update_AzureStorageSmbFileShare.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        endpointsUpdateAzureStorageSmbFileShare(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        Endpoint resource = manager.endpoints().getWithResponse("examples-rg", "examples-storageMoverName",
            "examples-endpointName", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }

    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/
     * Endpoints_Update_AzureStorageBlobContainer.json
     */
    /**
     * Sample code: Endpoints_Update_AzureStorageBlobContainer.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        endpointsUpdateAzureStorageBlobContainer(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        Endpoint resource = manager.endpoints().getWithResponse("examples-rg", "examples-storageMoverName",
            "examples-endpointName", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }

    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/
     * Endpoints_Update_SmbMount.json
     */
    /**
     * Sample code: Endpoints_Update_SmbMount.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsUpdateSmbMount(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        Endpoint resource = manager.endpoints().getWithResponse("examples-rg", "examples-storageMoverName",
            "examples-endpointName", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
