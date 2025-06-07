
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void operationsList(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
