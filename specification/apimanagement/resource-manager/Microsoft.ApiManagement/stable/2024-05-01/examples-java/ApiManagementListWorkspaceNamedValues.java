
/**
 * Samples for WorkspaceNamedValue ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceNamedValues.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceNamedValues.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceNamedValues(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().listByService("rg1", "apimService1", "wks1", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
