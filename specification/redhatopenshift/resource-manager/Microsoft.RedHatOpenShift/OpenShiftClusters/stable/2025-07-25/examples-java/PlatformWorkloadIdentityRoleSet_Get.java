
/**
 * Samples for PlatformWorkloadIdentityRoleSetOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-25/PlatformWorkloadIdentityRoleSet_Get.json
     */
    /**
     * Sample code: Gets a mapping of an OpenShift version to identity requirements, which includes operatorName,
     * roleDefinitionName, roleDefinitionId, and serviceAccounts.
     * 
     * @param manager Entry point to RedHatOpenShiftManager.
     */
    public static void
        getsAMappingOfAnOpenShiftVersionToIdentityRequirementsWhichIncludesOperatorNameRoleDefinitionNameRoleDefinitionIdAndServiceAccounts(
            com.azure.resourcemanager.redhatopenshift.RedHatOpenShiftManager manager) {
        manager.platformWorkloadIdentityRoleSetOperations().getWithResponse("location", "4.14",
            com.azure.core.util.Context.NONE);
    }
}
