
/**
 * Samples for WorkspaceDiscoveries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceDiscoveries_Get_InProgress.json
     */
    /**
     * Sample code: Read a running discovery with HTTP 200 without inventing accepted coverage.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void readARunningDiscoveryWithHTTP200WithoutInventingAcceptedCoverage(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceDiscoveries().getWithResponse("exampleRG", "exampleWorkspace",
            "00000000-0000-0000-0000-000000000009", com.azure.core.util.Context.NONE);
    }
}
