
/**
 * Samples for SyncGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/SyncGroupGetWithIdentity.json
     */
    /**
     * Sample code: Get a sync group with user assigned identity.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getASyncGroupWithUserAssignedIdentity(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncGroups().getWithResponse("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
