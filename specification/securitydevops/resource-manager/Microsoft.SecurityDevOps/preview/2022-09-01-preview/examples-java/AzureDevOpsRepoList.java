import com.azure.core.util.Context;

/** Samples for AzureDevOpsRepo List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoList.json
     */
    /**
     * Sample code: AzureDevOpsRepo_List.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsRepoList(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsRepoes().list("westusrg", "testconnector", "myOrg", "myProject", Context.NONE);
    }
}
