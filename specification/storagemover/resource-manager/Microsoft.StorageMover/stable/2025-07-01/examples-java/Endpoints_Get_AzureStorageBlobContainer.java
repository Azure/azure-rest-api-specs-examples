
/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Endpoints_Get_AzureStorageBlobContainer.json
     */
    /**
     * Sample code: Endpoints_Get_AzureStorageBlobContainer.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        endpointsGetAzureStorageBlobContainer(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.endpoints().getWithResponse("examples-rg", "examples-storageMoverName", "examples-endpointName",
            com.azure.core.util.Context.NONE);
    }
}
