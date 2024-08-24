
/**
 * Samples for ReplicationLinks Failover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ReplicationLinkFailover.json
     */
    /**
     * Sample code: Planned failover of a replication link.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void plannedFailoverOfAReplicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getReplicationLinks().failover("Default", "sourcesvr", "gamma-db",
            "4891ca10-ebd0-47d7-9182-c722651780fb", com.azure.core.util.Context.NONE);
    }
}
