import com.azure.core.util.Context;

/** Samples for AzureDevOpsConnector List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorListBySubscription.json
     */
    /**
     * Sample code: AzureDevOpsConnector_ListBySubscription.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsConnectorListBySubscription(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager.azureDevOpsConnectors().list(Context.NONE);
    }
}
