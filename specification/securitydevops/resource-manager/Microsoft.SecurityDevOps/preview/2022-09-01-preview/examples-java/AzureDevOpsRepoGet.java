import com.azure.core.util.Context;

/** Samples for AzureDevOpsRepo Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoGet.json
     */
    /**
     * Sample code: AzureDevOpsRepo_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsRepoGet(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager
            .azureDevOpsRepoes()
            .getWithResponse("westusrg", "testconnector", "myOrg", "myProject", "myRepo", Context.NONE);
    }
}
