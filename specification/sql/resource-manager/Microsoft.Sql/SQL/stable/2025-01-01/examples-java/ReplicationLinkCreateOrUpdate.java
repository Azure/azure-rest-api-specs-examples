
import com.azure.resourcemanager.sql.fluent.models.ReplicationLinkInner;
import com.azure.resourcemanager.sql.models.ReplicationLinkType;

/**
 * Samples for ReplicationLinks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ReplicationLinkCreateOrUpdate.json
     */
    /**
     * Sample code: Updates Replication Link.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updatesReplicationLink(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getReplicationLinks().createOrUpdate("Default", "sourcesvr", "gamma-db",
            "00000000-1111-2222-3333-666666666666",
            new ReplicationLinkInner().withLinkType(ReplicationLinkType.STANDBY), com.azure.core.util.Context.NONE);
    }
}
