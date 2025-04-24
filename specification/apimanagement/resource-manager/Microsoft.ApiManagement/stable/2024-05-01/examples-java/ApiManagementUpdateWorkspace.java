
import com.azure.resourcemanager.apimanagement.models.WorkspaceContract;

/**
 * Samples for Workspace Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspace.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspace.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspace(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        WorkspaceContract resource = manager.workspaces()
            .getWithResponse("rg1", "apimService1", "wks1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDisplayName("my workspace").withDescription("workspace 1").withIfMatch("*").apply();
    }
}
