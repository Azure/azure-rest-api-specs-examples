import com.azure.core.util.Context;

/** Samples for AzureDevOpsConnector GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorGet.json
     */
    /**
     * Sample code: AzureDevOpsConnector_Get.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsConnectorGet(com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsConnectors().getByResourceGroupWithResponse("westusrg", "testconnector", Context.NONE);
    }
}
