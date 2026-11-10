
/**
 * Samples for WorkspaceDiscoveries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceDiscoveries_Get_BaselineUnknown.json
     */
    /**
     * Sample code: Get a healthy container scan with a structurally unknown inventory denominator.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAHealthyContainerScanWithAStructurallyUnknownInventoryDenominator(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceDiscoveries().getWithResponse("exampleRG", "exampleWorkspace",
            "00000000-0000-0000-0000-000000000006", com.azure.core.util.Context.NONE);
    }
}
