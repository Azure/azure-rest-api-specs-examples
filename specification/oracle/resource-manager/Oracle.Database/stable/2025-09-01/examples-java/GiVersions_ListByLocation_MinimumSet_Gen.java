
/**
 * Samples for GiVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/GiVersions_ListByLocation_MinimumSet_Gen.json
     */
    /**
     * Sample code: GiVersions_ListByLocation_MinimumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        giVersionsListByLocationMinimumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.giVersions().listByLocation("eastus", null, null, null, com.azure.core.util.Context.NONE);
    }
}
