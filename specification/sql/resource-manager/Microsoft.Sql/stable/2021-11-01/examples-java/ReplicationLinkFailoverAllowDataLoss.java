
/**
 * Samples for ReplicationLinks FailoverAllowDataLoss.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ReplicationLinkFailoverAllowDataLoss.
     * json
     */
    /**
     * Sample code: Forced failover of a replication link.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void forcedFailoverOfAReplicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getReplicationLinks().failoverAllowDataLoss("Default", "sourcesvr",
            "gamma-db", "4891ca10-ebd0-47d7-9182-c722651780fb", com.azure.core.util.Context.NONE);
    }
}
