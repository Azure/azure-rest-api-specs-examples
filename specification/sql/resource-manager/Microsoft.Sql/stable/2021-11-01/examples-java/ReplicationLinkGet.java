
/**
 * Samples for ReplicationLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ReplicationLinkGet.json
     */
    /**
     * Sample code: Gets the replication link.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheReplicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getReplicationLinks().getWithResponse("Default", "sourcesvr",
            "gamma-db", "4891ca10-ebd0-47d7-9182-c722651780fb", com.azure.core.util.Context.NONE);
    }
}
