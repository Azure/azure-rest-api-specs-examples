
/**
 * Samples for NetworkAnchors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkAnchors_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkAnchors_Delete_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        networkAnchorsDeleteMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.networkAnchors().delete("rgopenapi", "networkAnchor1", com.azure.core.util.Context.NONE);
    }
}
