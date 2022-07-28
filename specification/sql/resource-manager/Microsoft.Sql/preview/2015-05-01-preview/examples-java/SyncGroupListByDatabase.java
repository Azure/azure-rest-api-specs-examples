import com.azure.core.util.Context;

/** Samples for SyncGroups ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/examples/SyncGroupListByDatabase.json
     */
    /**
     * Sample code: List sync groups under a given database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSyncGroupsUnderAGivenDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getSyncGroups()
            .listByDatabase("syncgroupcrud-65440", "syncgroupcrud-8475", "syncgroupcrud-4328", Context.NONE);
    }
}
