
/**
 * Samples for ReplicationLinks ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ReplicationLinkListByServer.json
     */
    /**
     * Sample code: List replication links on server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listReplicationLinksOnServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getReplicationLinks().listByServer("Default", "sourcesvr",
            com.azure.core.util.Context.NONE);
    }
}
