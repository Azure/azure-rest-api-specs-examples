
/**
 * Samples for WorkspaceTag Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceTag.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceTag.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTags().getWithResponse("rg1", "apimService1", "wks1", "59306a29e4bbd510dc24e5f9",
            com.azure.core.util.Context.NONE);
    }
}
