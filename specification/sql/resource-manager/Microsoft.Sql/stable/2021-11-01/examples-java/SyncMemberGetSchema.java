
/**
 * Samples for SyncMembers ListMemberSchemas.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncMemberGetSchema.json
     */
    /**
     * Sample code: Get a sync member schema.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASyncMemberSchema(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncMembers().listMemberSchemas("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", "syncgroupcrud-4879",
            com.azure.core.util.Context.NONE);
    }
}
