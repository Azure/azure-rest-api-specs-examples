
/**
 * Samples for SovereignViewOperations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-31-preview/SovereignViewOperations_Delete.json
     */
    /**
     * Sample code: SovereignViewOperations_Delete.
     * 
     * @param manager Entry point to SovereignManager.
     */
    public static void sovereignViewOperationsDelete(com.azure.resourcemanager.sovereign.SovereignManager manager) {
        manager.sovereignViewOperations().delete("ExampleResourceGroup", "ExampleSV", com.azure.core.util.Context.NONE);
    }
}
