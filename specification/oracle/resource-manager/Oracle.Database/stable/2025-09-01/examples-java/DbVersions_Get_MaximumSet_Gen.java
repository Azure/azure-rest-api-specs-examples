
/**
 * Samples for DbVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/DbVersions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DbVersions_Get_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dbVersionsGetMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbVersions().getWithResponse("eastus", "23.0.0.0.0", com.azure.core.util.Context.NONE);
    }
}
