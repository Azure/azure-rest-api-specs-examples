import com.azure.core.util.Context;

/** Samples for AzureDevOpsOrg List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsOrgList.json
     */
    /**
     * Sample code: AzureDevOpsOrg_List.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsOrgList(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsOrgs().list("westusrg", "testconnector", Context.NONE);
    }
}
