
/**
 * Samples for ResourceAnchors GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ResourceAnchors_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ResourceAnchors_Get_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        resourceAnchorsGetMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.resourceAnchors().getByResourceGroupWithResponse("rgopenapi", "resourceanchor1",
            com.azure.core.util.Context.NONE);
    }
}
