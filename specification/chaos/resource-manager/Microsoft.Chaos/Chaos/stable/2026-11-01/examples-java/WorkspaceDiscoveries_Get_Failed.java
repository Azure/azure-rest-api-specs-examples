
/**
 * Samples for WorkspaceDiscoveries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceDiscoveries_Get_Failed.json
     */
    /**
     * Sample code: Get a terminal failed discovery without an accepted inventory result.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getATerminalFailedDiscoveryWithoutAnAcceptedInventoryResult(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceDiscoveries().getWithResponse("exampleRG", "exampleWorkspace",
            "00000000-0000-0000-0000-000000000005", com.azure.core.util.Context.NONE);
    }
}
