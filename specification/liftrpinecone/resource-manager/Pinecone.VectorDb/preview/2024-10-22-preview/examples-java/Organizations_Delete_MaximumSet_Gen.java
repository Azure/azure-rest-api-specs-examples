
/**
 * Samples for Organizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-22-preview/Organizations_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Delete_MaximumSet.
     * 
     * @param manager Entry point to PineconeVectorDbManager.
     */
    public static void
        organizationsDeleteMaximumSet(com.azure.resourcemanager.pineconevectordb.PineconeVectorDbManager manager) {
        manager.organizations().delete("rgopenapi", "example-organization-name", com.azure.core.util.Context.NONE);
    }
}
