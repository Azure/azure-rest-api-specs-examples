
import com.azure.resourcemanager.oracledatabase.models.PeerDbDetails;

/**
 * Samples for AutonomousDatabases Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabase_failover.json
     */
    /**
     * Sample code: AutonomousDatabases_Failover.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesFailover(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().failover("rg000", "databasedb1", new PeerDbDetails().withPeerDbId("peerDbId"),
            com.azure.core.util.Context.NONE);
    }
}
