
/**
 * Samples for PrivateLinkResource Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/PrivateLinkResourceGet.json
     */
    /**
     * Sample code: GetPrivateLinkResource.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void getPrivateLinkResource(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.privateLinkResources().getWithResponse("default-azurebatch-japaneast", "sampleacct", "batchAccount",
            com.azure.core.util.Context.NONE);
    }
}
