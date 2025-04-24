
/**
 * Samples for WorkspaceApi Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceApiContract.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceApiContract.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceApiContract(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApis().getWithResponse("rg1", "apimService1", "wks1", "57d1f7558aa04f15146d9d8a",
            com.azure.core.util.Context.NONE);
    }
}
