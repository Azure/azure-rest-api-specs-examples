
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void operationsListMaximumSet(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
