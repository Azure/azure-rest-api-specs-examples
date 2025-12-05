
/**
 * Samples for NetAppResourceRegionInfos List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/RegionInfos_List.json
     */
    /**
     * Sample code: RegionInfos_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void regionInfosList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceRegionInfos().list("eastus", com.azure.core.util.Context.NONE);
    }
}
