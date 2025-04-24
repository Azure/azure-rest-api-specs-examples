
/**
 * Samples for WorkspaceGroupUser Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceGroupUser.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceGroupUser.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceGroupUser(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGroupUsers().createWithResponse("rg1", "apimService1", "wks1", "tempgroup",
            "59307d350af58404d8a26300", com.azure.core.util.Context.NONE);
    }
}
