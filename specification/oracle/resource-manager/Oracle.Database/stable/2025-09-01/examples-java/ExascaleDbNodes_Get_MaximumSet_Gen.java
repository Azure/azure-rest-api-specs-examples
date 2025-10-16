
/**
 * Samples for ExascaleDbNodes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExascaleDbNodes_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbNodes_Get_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exascaleDbNodesGetMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbNodes().getWithResponse("rgopenapi", "exadbvmcluster1", "exascaledbnode1",
            com.azure.core.util.Context.NONE);
    }
}
