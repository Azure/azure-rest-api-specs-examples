import com.azure.core.util.Context;

/** Samples for AzureDevOpsProject List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsProjectList.json
     */
    /**
     * Sample code: AzureDevOpsProject_List.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsProjectList(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsProjects().list("westusrg", "testconnector", "myOrg", Context.NONE);
    }
}
