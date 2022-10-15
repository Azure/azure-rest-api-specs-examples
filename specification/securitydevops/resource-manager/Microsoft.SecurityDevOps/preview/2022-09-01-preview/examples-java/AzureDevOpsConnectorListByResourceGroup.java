import com.azure.core.util.Context;

/** Samples for AzureDevOpsConnector ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorListByResourceGroup.json
     */
    /**
     * Sample code: AzureDevOpsConnector_ListByResourceGroup.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsConnectorListByResourceGroup(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsConnectors().listByResourceGroup("westusrg", Context.NONE);
    }
}
