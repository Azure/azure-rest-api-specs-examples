
/**
 * Samples for SyncGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/SyncGroupGet.json
     */
    /**
     * Sample code: Get a sync group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getASyncGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSyncGroups().getWithResponse("syncgroupcrud-65440",
            "syncgroupcrud-8475", "syncgroupcrud-4328", "syncgroupcrud-3187", com.azure.core.util.Context.NONE);
    }
}
