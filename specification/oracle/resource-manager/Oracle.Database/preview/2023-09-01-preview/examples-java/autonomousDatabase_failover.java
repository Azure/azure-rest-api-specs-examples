
import com.azure.resourcemanager.oracledatabase.models.PeerDbDetails;

/**
 * Samples for AutonomousDatabases Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * autonomousDatabase_failover.json
     */
    /**
     * Sample code: Perform failover action on Autonomous Database.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void performFailoverActionOnAutonomousDatabase(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().failover("rg000", "databasedb1", new PeerDbDetails().withPeerDbId("peerDbId"),
            com.azure.core.util.Context.NONE);
    }
}
