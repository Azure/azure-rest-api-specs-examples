
/**
 * Samples for WorkspaceApiSchema Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceApiSchema.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceApiSchema.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceApiSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiSchemas().getWithResponse("rg1", "apimService1", "wks1", "59d6bb8f1f7fab13dc67ec9b",
            "ec12520d-9d48-4e7b-8f39-698ca2ac63f1", com.azure.core.util.Context.NONE);
    }
}
