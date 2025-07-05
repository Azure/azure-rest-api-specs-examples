
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void operationsListMinimumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
