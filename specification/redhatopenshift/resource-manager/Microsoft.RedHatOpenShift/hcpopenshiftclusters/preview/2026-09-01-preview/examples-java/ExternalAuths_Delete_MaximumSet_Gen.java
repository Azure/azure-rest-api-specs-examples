
/**
 * Samples for ExternalAuths Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/ExternalAuths_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExternalAuths_Delete_MaximumSet.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void externalAuthsDeleteMaximumSet(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.externalAuths().delete("rgopenapi", "hcpCluster-name", "my-cool-auth",
            com.azure.core.util.Context.NONE);
    }
}
