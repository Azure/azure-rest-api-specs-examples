
/**
 * Samples for SyncMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/SyncMemberGetWithIdentity.json
     */
    /**
     * Sample code: Get a sync member with user assigned identity.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getASyncMemberWithUserAssignedIdentity(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncMembers().getWithResponse("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", "syncmembercrud-4879", com.azure.core.util.Context.NONE);
    }
}
