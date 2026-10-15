
/**
 * Samples for EdgeActionVersions SwapDefault.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeActionVersions_SwapDefault.json
     */
    /**
     * Sample code: Swap Default Version.
     * 
     * @param manager Entry point to EdgeActionsManager.
     */
    public static void swapDefaultVersion(com.azure.resourcemanager.edgeactions.EdgeActionsManager manager) {
        manager.edgeActionVersions().swapDefault("testrg", "edgeAction1", "version1", com.azure.core.util.Context.NONE);
    }
}
