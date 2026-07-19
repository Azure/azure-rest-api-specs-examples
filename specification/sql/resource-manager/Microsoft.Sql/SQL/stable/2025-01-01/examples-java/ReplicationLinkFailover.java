
/**
 * Samples for ReplicationLinks Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ReplicationLinkFailover.json
     */
    /**
     * Sample code: Planned failover of a replication link.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void plannedFailoverOfAReplicationLink(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getReplicationLinks().failover("Default", "sourcesvr", "gamma-db",
            "4891ca10-ebd0-47d7-9182-c722651780fb", com.azure.core.util.Context.NONE);
    }
}
