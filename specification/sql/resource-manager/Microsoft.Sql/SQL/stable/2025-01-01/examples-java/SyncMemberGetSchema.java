
/**
 * Samples for SyncMembers ListMemberSchemas.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SyncMemberGetSchema.json
     */
    /**
     * Sample code: Get a sync member schema.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getASyncMemberSchema(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSyncMembers().listMemberSchemas("syncgroupcrud-65440", "syncgroupcrud-8475",
            "syncgroupcrud-4328", "syncgroupcrud-3187", "syncgroupcrud-4879", com.azure.core.util.Context.NONE);
    }
}
