
import com.azure.resourcemanager.sql.models.ReplicaType;

/**
 * Samples for Databases Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FailoverDatabase.json
     */
    /**
     * Sample code: Failover an database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void failoverAnDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().failover("group1", "testServer", "testDatabase", ReplicaType.PRIMARY,
            com.azure.core.util.Context.NONE);
    }
}
