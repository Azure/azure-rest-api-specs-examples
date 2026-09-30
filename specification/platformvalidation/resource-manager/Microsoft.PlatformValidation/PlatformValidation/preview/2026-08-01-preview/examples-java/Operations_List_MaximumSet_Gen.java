
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet.
     * 
     * @param manager Entry point to PlatformValidationManager.
     */
    public static void
        operationsListMaximumSet(com.azure.resourcemanager.platformvalidation.PlatformValidationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
