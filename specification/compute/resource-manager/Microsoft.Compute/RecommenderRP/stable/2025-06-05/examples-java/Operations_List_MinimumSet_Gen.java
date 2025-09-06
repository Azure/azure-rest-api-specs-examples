
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-05/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeRecommenderManager.
     */
    public static void
        operationsListMinimumSetGen(com.azure.resourcemanager.computerecommender.ComputeRecommenderManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
