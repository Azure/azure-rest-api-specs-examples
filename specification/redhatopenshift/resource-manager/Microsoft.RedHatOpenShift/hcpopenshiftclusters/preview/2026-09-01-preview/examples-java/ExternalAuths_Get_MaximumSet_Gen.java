
/**
 * Samples for ExternalAuths Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/ExternalAuths_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExternalAuths_Get_MaximumSet.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void externalAuthsGetMaximumSet(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.externalAuths().getWithResponse("rgopenapi", "hcpCluster-name", "my-cool-auth",
            com.azure.core.util.Context.NONE);
    }
}
