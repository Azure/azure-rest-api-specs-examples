
import com.azure.resourcemanager.oracledatabase.models.DisasterRecoveryConfigurationDetails;
import com.azure.resourcemanager.oracledatabase.models.DisasterRecoveryType;

/**
 * Samples for AutonomousDatabases ChangeDisasterRecoveryConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabase_changeDisasterRecoveryConfiguration.json
     */
    /**
     * Sample code: AutonomousDatabases_ChangeDisasterRecoveryConfiguration.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void autonomousDatabasesChangeDisasterRecoveryConfiguration(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().changeDisasterRecoveryConfiguration(
            "rg000", "databasedb1", new DisasterRecoveryConfigurationDetails()
                .withDisasterRecoveryType(DisasterRecoveryType.ADG).withIsReplicateAutomaticBackups(false),
            com.azure.core.util.Context.NONE);
    }
}
