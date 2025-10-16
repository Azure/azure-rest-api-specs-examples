
/**
 * Samples for ResourceAnchors ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ResourceAnchors_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: ResourceAnchors_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void resourceAnchorsListByResourceGroupMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.resourceAnchors().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
