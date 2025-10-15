
/**
 * Samples for ResourceAnchors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ResourceAnchors_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ResourceAnchors_Delete_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        resourceAnchorsDeleteMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.resourceAnchors().delete("rgopenapi", "resourceanchor1", com.azure.core.util.Context.NONE);
    }
}
