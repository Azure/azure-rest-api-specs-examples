
/**
 * Samples for ApiManagementWorkspaceLinks ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceLinks.json
     */
    /**
     * Sample code: ApiManagementListGatewayConfigConnection.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListGatewayConfigConnection(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementWorkspaceLinks().listByService("rg1", "service1", com.azure.core.util.Context.NONE);
    }
}
