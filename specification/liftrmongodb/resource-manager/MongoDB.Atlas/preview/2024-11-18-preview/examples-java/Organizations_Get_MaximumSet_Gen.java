
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-18-preview/Organizations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Get_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void organizationsGetMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.organizations().getByResourceGroupWithResponse("rgopenapi", "U.1-:7", com.azure.core.util.Context.NONE);
    }
}
