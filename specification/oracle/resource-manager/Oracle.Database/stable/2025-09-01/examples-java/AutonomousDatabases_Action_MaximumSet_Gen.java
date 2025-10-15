
import com.azure.resourcemanager.oracledatabase.models.AutonomousDatabaseLifecycleAction;
import com.azure.resourcemanager.oracledatabase.models.AutonomousDatabaseLifecycleActionEnum;

/**
 * Samples for AutonomousDatabases Action.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/AutonomousDatabases_Action_MaximumSet_Gen.json
     */
    /**
     * Sample code: AutonomousDatabases_Action_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesActionMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().action("rgopenapi", "databasedb1",
            new AutonomousDatabaseLifecycleAction().withAction(AutonomousDatabaseLifecycleActionEnum.START),
            com.azure.core.util.Context.NONE);
    }
}
