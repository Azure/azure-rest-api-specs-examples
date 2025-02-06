
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-22-preview/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumSet.
     * 
     * @param manager Entry point to PineconeVectorDbManager.
     */
    public static void
        operationsListMinimumSet(com.azure.resourcemanager.pineconevectordb.PineconeVectorDbManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
