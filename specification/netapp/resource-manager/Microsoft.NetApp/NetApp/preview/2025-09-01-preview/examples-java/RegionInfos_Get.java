
/**
 * Samples for NetAppResourceRegionInfos Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/RegionInfos_Get.json
     */
    /**
     * Sample code: RegionInfos_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void regionInfosGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceRegionInfos().getWithResponse("eastus", com.azure.core.util.Context.NONE);
    }
}
