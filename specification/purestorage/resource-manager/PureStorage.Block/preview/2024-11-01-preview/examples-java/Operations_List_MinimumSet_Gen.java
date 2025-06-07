
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Operations_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MinimumGen_Set.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        operationsListMinimumGenSet(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
