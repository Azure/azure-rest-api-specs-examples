
/**
 * Samples for HcpOperatorIdentityRoleSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01-preview/HcpOperatorIdentityRoleSets_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: HcpOperatorIdentityRoleSets_Get_MaximumSet.
     * 
     * @param manager Entry point to RedHatOpenShiftHostedControlPlanesManager.
     */
    public static void hcpOperatorIdentityRoleSetsGetMaximumSet(
        com.azure.resourcemanager.redhatopenshifthcp.RedHatOpenShiftHostedControlPlanesManager manager) {
        manager.hcpOperatorIdentityRoleSets().getWithResponse("uksouth", "hcp-example-role-set",
            com.azure.core.util.Context.NONE);
    }
}
