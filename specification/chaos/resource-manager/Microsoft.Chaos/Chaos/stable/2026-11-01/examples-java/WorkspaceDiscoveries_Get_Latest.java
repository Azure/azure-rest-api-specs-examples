
/**
 * Samples for WorkspaceDiscoveries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceDiscoveries_Get_Latest.json
     */
    /**
     * Sample code: Read latest as an alias while returning the canonical discovery identity.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void readLatestAsAnAliasWhileReturningTheCanonicalDiscoveryIdentity(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceDiscoveries().getWithResponse("exampleRG", "exampleWorkspace", "latest",
            com.azure.core.util.Context.NONE);
    }
}
