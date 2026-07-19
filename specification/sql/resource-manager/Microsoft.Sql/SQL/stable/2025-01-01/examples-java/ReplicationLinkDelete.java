
/**
 * Samples for ReplicationLinks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ReplicationLinkDelete.json
     */
    /**
     * Sample code: Delete replication link on server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteReplicationLinkOnServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getReplicationLinks().delete("Default", "sourcesvr", "gamma-db",
            "4891ca10-ebd0-47d7-9182-c722651780fb", com.azure.core.util.Context.NONE);
    }
}
