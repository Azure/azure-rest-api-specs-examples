
/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/
     * Endpoints_Get_SmbMount.json
     */
    /**
     * Sample code: Endpoints_Get_SmbMount.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsGetSmbMount(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.endpoints().getWithResponse("examples-rg", "examples-storageMoverName", "examples-endpointName",
            com.azure.core.util.Context.NONE);
    }
}
