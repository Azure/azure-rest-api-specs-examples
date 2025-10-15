
/**
 * Samples for DbSystems GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/DbSystems_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DbSystems_Get_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dbSystemsGetMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbSystems().getByResourceGroupWithResponse("rgopenapi", "dbsystem1", com.azure.core.util.Context.NONE);
    }
}
