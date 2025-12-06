
/**
 * Samples for ElasticVolumes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticVolumes_Get.json
     */
    /**
     * Sample code: ElasticVolumes_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticVolumesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticVolumes().getWithResponse("myRG", "account1", "pool1", "volume1",
            com.azure.core.util.Context.NONE);
    }
}
