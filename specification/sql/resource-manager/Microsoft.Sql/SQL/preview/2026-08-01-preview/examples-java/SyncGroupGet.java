
/**
 * Samples for SyncGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/SyncGroupGet.json
     */
    /**
     * Sample code: Get a sync group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getASyncGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().getWithResponse("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
