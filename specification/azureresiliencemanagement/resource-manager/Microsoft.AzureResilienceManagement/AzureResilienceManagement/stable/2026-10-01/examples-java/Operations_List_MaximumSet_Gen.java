
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        operationsListMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
