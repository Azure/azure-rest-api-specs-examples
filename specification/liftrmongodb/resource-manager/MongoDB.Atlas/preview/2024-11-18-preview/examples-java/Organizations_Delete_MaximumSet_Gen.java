
/**
 * Samples for Organizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-18-preview/Organizations_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Delete_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void
        organizationsDeleteMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.organizations().delete("rgopenapi", "U.1-:7", com.azure.core.util.Context.NONE);
    }
}
