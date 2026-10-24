
/**
 * Samples for HcpOperatorIdentityRoleSets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/HcpOperatorIdentityRoleSets_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: HcpOperatorIdentityRoleSets_List_MaximumSet.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void hcpOperatorIdentityRoleSetsListMaximumSet(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.hcpOperatorIdentityRoleSets().list("uksouth", com.azure.core.util.Context.NONE);
    }
}
