
/**
 * Samples for ExascaleDbNodes ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExascaleDbNodes_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbNodes_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exascaleDbNodesListByParentMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbNodes().listByParent("rgopenapi", "vmcluster", com.azure.core.util.Context.NONE);
    }
}
