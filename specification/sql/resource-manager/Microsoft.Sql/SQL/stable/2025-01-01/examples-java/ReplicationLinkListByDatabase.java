
/**
 * Samples for ReplicationLinks ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ReplicationLinkListByDatabase.json
     */
    /**
     * Sample code: List replication links on server on database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listReplicationLinksOnServerOnDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getReplicationLinks().listByDatabase("Default", "sourcesvr", "tetha-db",
            com.azure.core.util.Context.NONE);
    }
}
