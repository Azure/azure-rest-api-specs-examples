import com.azure.core.util.Context;

/** Samples for AzureDevOpsRepo ListByConnector. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoListByConnector.json
     */
    /**
     * Sample code: AzureDevOpsRepo_ListByConnector.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsRepoListByConnector(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsRepoes().listByConnector("westusrg", "testconnector", Context.NONE);
    }
}
