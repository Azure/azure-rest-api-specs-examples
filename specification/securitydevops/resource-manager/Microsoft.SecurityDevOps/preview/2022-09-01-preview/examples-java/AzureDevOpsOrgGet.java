import com.azure.core.util.Context;

/** Samples for AzureDevOpsOrg Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsOrgGet.json
     */
    /**
     * Sample code: AzureDevOpsOrg_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsOrgGet(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsOrgs().getWithResponse("westusrg", "testconnector", "myOrg", Context.NONE);
    }
}
