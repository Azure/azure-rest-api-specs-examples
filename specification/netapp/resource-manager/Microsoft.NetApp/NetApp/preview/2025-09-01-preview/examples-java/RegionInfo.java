
/**
 * Samples for NetAppResource QueryRegionInfo.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/RegionInfo.json
     */
    /**
     * Sample code: RegionInfo_Query.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void regionInfoQuery(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResources().queryRegionInfoWithResponse("eastus", com.azure.core.util.Context.NONE);
    }
}
