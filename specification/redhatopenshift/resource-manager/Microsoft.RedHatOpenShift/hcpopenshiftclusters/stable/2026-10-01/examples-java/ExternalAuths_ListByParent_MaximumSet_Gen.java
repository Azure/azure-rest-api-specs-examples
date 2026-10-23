
/**
 * Samples for ExternalAuths ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ExternalAuths_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExternalAuths_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void externalAuthsListByParentMaximumSet(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.externalAuths().listByParent("rgopenapi", "hcpCluster-name", com.azure.core.util.Context.NONE);
    }
}
