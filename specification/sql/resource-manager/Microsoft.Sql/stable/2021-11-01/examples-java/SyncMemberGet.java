
/**
 * Samples for SyncMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncMemberGet.json
     */
    /**
     * Sample code: Get a sync member.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASyncMember(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncMembers().getWithResponse("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", "syncmembercrud-4879",
            com.azure.core.util.Context.NONE);
    }
}
