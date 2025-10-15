
/**
 * Samples for NetworkAnchors ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkAnchors_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkAnchors_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void networkAnchorsListByResourceGroupMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.networkAnchors().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
