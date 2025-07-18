
import com.azure.resourcemanager.oracledatabase.models.PeerDbDetails;

/**
 * Samples for AutonomousDatabases Switchover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/autonomousDatabase_switchover.json
     */
    /**
     * Sample code: AutonomousDatabases_Switchover.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesSwitchover(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().switchover("rg000", "databasedb1", new PeerDbDetails().withPeerDbId("peerDbId"),
            com.azure.core.util.Context.NONE);
    }
}
