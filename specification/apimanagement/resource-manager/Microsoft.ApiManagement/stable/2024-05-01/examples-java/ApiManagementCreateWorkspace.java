
/**
 * Samples for Workspace CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspace.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspace.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspace(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaces().define("wks1").withExistingService("rg1", "apimService1").withDisplayName("my workspace")
            .withDescription("workspace 1").create();
    }
}
