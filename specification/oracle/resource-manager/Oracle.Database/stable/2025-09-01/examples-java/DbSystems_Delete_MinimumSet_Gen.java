
/**
 * Samples for DbSystems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/DbSystems_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: DbSystems_Delete_MinimumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        dbSystemsDeleteMinimumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbSystems().delete("rgopenapi", "dbsystem1", com.azure.core.util.Context.NONE);
    }
}
