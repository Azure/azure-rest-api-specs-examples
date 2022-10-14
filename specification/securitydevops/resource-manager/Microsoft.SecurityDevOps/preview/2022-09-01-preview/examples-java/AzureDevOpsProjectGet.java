import com.azure.core.util.Context;

/** Samples for AzureDevOpsProject Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsProjectGet.json
     */
    /**
     * Sample code: AzureDevOpsProject_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsProjectGet(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsProjects().getWithResponse("westusrg", "testconnector", "myOrg", "myProject", Context.NONE);
    }
}
