
/**
 * Samples for WorkspaceApiSchema GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceApiSchema.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceApiSchema.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceApiSchema(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiSchemas().getEntityTagWithResponse("rg1", "apimService1", "wks1",
            "57d1f7558aa04f15146d9d8a", "ec12520d-9d48-4e7b-8f39-698ca2ac63f1", com.azure.core.util.Context.NONE);
    }
}
