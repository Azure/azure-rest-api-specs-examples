
/**
 * Samples for EdgeActionVersions SwapDefault.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/EdgeActionVersions_SwapDefault.json
     */
    /**
     * Sample code: Swap Default Version.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void swapDefaultVersion(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().swapDefault("testrg", "edgeAction1", "1.0", com.azure.core.util.Context.NONE);
    }
}
