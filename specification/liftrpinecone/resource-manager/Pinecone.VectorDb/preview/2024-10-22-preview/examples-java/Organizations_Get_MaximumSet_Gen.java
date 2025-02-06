
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-22-preview/Organizations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Get_MaximumSet.
     * 
     * @param manager Entry point to PineconeVectorDbManager.
     */
    public static void
        organizationsGetMaximumSet(com.azure.resourcemanager.pineconevectordb.PineconeVectorDbManager manager) {
        manager.organizations().getByResourceGroupWithResponse("rgopenapi", "example-organization-name",
            com.azure.core.util.Context.NONE);
    }
}
