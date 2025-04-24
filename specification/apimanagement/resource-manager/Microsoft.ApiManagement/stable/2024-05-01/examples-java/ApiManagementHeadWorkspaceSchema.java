
/**
 * Samples for WorkspaceGlobalSchema GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceSchema.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceSchema.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGlobalSchemas().getEntityTagWithResponse("rg1", "apimService1", "wks1", "myschema",
            com.azure.core.util.Context.NONE);
    }
}
