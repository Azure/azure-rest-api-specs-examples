
/**
 * Samples for NetAppResourceRegionInfos Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/RegionInfos_Get.json
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
