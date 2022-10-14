import com.azure.core.util.Context;

/** Samples for AzureDevOpsConnectorStatsOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorStatsGet.json
     */
    /**
     * Sample code: AzureDevOpsConnectorStats_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsConnectorStatsGet(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsConnectorStatsOperations().getWithResponse("westusrg", "testconnector", Context.NONE);
    }
}
