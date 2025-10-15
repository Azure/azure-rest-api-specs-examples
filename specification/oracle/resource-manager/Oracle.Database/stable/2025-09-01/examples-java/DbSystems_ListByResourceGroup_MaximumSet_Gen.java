
/**
 * Samples for DbSystems ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/DbSystems_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DbSystems_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        dbSystemsListByResourceGroupMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbSystems().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
