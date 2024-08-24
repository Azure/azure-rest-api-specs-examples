
/**
 * Samples for ReplicationLinks ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ReplicationLinkListByDatabase.json
     */
    /**
     * Sample code: List replication links on server on database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listReplicationLinksOnServerOnDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getReplicationLinks().listByDatabase("Default", "sourcesvr",
            "tetha-db", com.azure.core.util.Context.NONE);
    }
}
