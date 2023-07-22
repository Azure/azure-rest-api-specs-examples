/** Samples for Endpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Endpoints_Get_AzureStorageSmbFileShare.json
     */
    /**
     * Sample code: Endpoints_Get_AzureStorageSmbFileShare.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsGetAzureStorageSmbFileShare(
        com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .endpoints()
            .getWithResponse(
                "examples-rg", "examples-storageMoverName", "examples-endpointName", com.azure.core.util.Context.NONE);
    }
}
