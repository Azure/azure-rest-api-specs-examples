
/**
 * Samples for ReplicationLinks ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ReplicationLinkListByServer.json
     */
    /**
     * Sample code: List replication links on server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listReplicationLinksOnServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getReplicationLinks().listByServer("Default", "sourcesvr",
            com.azure.core.util.Context.NONE);
    }
}
