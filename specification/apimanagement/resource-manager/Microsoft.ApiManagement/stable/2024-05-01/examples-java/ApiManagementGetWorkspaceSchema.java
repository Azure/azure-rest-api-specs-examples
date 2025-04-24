
/**
 * Samples for WorkspaceGlobalSchema Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceSchema.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceSchema.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGlobalSchemas().getWithResponse("rg1", "apimService1", "wks1", "schema1",
            com.azure.core.util.Context.NONE);
    }
}
