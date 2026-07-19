
import com.azure.resourcemanager.sql.models.ReplicationLinkType;
import com.azure.resourcemanager.sql.models.ReplicationLinkUpdate;

/**
 * Samples for ReplicationLinks Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ReplicationLinkUpdate.json
     */
    /**
     * Sample code: Update Replication Link.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateReplicationLink(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getReplicationLinks().update("Default", "sourcesvr", "gamma-db",
            "00000000-1111-2222-3333-666666666666",
            new ReplicationLinkUpdate().withLinkType(ReplicationLinkType.STANDBY), com.azure.core.util.Context.NONE);
    }
}
