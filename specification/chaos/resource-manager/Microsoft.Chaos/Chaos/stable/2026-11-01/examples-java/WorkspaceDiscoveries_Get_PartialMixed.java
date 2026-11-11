
/**
 * Samples for WorkspaceDiscoveries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceDiscoveries_Get_PartialMixed.json
     */
    /**
     * Sample code: Get a partially succeeded result with classified guidance, system uncertainty, and a truncated
     * warning sample.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAPartiallySucceededResultWithClassifiedGuidanceSystemUncertaintyAndATruncatedWarningSample(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceDiscoveries().getWithResponse("exampleRG", "exampleWorkspace",
            "00000000-0000-0000-0000-000000000008", com.azure.core.util.Context.NONE);
    }
}
