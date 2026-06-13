
/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/Endpoints_Get_S3WithHMAC.json
     */
    /**
     * Sample code: Endpoints_Get_S3WithHmac.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsGetS3WithHmac(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.endpoints().getWithResponse("examples-rg", "examples-storageMoverName", "examples-endpointName",
            com.azure.core.util.Context.NONE);
    }
}
