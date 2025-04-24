
/**
 * Samples for WorkspaceNamedValue Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceNamedValueWithKeyVault.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceNamedValueWithKeyVault.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetWorkspaceNamedValueWithKeyVault(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().getWithResponse("rg1", "apimService1", "wks1", "testprop6",
            com.azure.core.util.Context.NONE);
    }
}
