
/**
 * Samples for SyncMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncMemberGet.json
     */
    /**
     * Sample code: Get a sync member.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getASyncMember(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncMembers().getWithResponse("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", "syncmembercrud-4879", com.azure.core.util.Context.NONE);
    }
}
