
/**
 * Samples for SovereignViewOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-31-preview/SovereignViewOperations_ListBySubscription.json
     */
    /**
     * Sample code: SovereignViewOperations_ListBySubscription.
     * 
     * @param manager Entry point to SovereignManager.
     */
    public static void
        sovereignViewOperationsListBySubscription(com.azure.resourcemanager.sovereign.SovereignManager manager) {
        manager.sovereignViewOperations().list(com.azure.core.util.Context.NONE);
    }
}
