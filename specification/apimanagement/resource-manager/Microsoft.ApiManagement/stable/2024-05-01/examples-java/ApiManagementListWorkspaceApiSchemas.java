
/**
 * Samples for WorkspaceApiSchema ListByApi.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListWorkspaceApiSchemas.json
     */
    /**
     * Sample code: ApiManagementListWorkspaceApiSchemas.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListWorkspaceApiSchemas(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiSchemas().listByApi("rg1", "apimService1", "wks1", "59d5b28d1f7fab116c282650", null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
