import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.UnlinkParameters;

/** Samples for ReplicationLinks Unlink. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ReplicationLinkUnlink.json
     */
    /**
     * Sample code: Delete replication link.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteReplicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getReplicationLinks()
            .unlink(
                "sqlcrudtest-8931",
                "sqlcrudtest-2137",
                "testdb",
                "f0550bf5-07ce-4270-8e4b-71737975973a",
                new UnlinkParameters().withForcedTermination(true),
                Context.NONE);
    }
}
