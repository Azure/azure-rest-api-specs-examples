
/**
 * Samples for WorkspaceDiscoveries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceDiscoveries_Get_Established.json
     */
    /**
     * Sample code: Get an established discovery result with a known resource population.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAnEstablishedDiscoveryResultWithAKnownResourcePopulation(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceDiscoveries().getWithResponse("exampleRG", "exampleWorkspace",
            "00000000-0000-0000-0000-000000000002", com.azure.core.util.Context.NONE);
    }
}
