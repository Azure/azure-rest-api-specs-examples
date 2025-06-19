
/**
 * Samples for GiMinorVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/GiMinorVersions_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: GiMinorVersions_Get_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        giMinorVersionsGetMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.giMinorVersions().getWithResponse("eastus", "giVersionName", "giMinorVersionName",
            com.azure.core.util.Context.NONE);
    }
}
