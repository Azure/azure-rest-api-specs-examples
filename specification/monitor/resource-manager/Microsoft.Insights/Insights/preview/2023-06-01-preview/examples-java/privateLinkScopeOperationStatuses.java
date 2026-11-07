
/**
 * Samples for PrivateLinkScopeOperationStatus GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-06-01-preview/privateLinkScopeOperationStatuses.json
     */
    /**
     * Sample code: Get specific operation status.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getSpecificOperationStatus(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getPrivateLinkScopeOperationStatus().getByResourceGroupWithResponse("MyResourceGroup",
            "713192d7-503f-477a-9cfe-4efc3ee2bd11", com.azure.core.util.Context.NONE);
    }
}
