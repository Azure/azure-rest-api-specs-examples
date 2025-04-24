
/**
 * Samples for WorkspaceNamedValue Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceNamedValue.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceNamedValue.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceNamedValue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().getWithResponse("rg1", "apimService1", "wks1", "testarmTemplateproperties2",
            com.azure.core.util.Context.NONE);
    }
}
