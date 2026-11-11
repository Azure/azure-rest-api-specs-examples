
/**
 * Samples for SovereignViewOperations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-31-preview/SovereignViewOperations_ListByResourceGroup.json
     */
    /**
     * Sample code: SovereignViewOperations_ListByResourceGroup.
     * 
     * @param manager Entry point to SovereignManager.
     */
    public static void
        sovereignViewOperationsListByResourceGroup(com.azure.resourcemanager.sovereign.SovereignManager manager) {
        manager.sovereignViewOperations().listByResourceGroup("ExampleResourceGroup", com.azure.core.util.Context.NONE);
    }
}
