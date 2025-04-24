
/**
 * Samples for WorkspaceNamedValue GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceNamedValue.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceNamedValue.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceNamedValue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().getEntityTagWithResponse("rg1", "apimService1", "wks1",
            "testarmTemplateproperties2", com.azure.core.util.Context.NONE);
    }
}
