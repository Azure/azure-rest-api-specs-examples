
/**
 * Samples for PrivateLinkScopeOperationStatus GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/
     * privateLinkScopeOperationStatuses.json
     */
    /**
     * Sample code: Get specific operation status.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSpecificOperationStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getPrivateLinkScopeOperationStatus()
            .getByResourceGroupWithResponse("MyResourceGroup", "713192d7-503f-477a-9cfe-4efc3ee2bd11",
                com.azure.core.util.Context.NONE);
    }
}
