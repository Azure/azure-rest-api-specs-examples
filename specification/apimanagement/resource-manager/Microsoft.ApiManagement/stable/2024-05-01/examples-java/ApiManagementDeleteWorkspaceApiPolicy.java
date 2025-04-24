
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceApiPolicy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceApiPolicy.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceApiPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspaceApiPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiPolicies().deleteWithResponse("rg1", "apimService1", "wks1", "loggerId",
            PolicyIdName.POLICY, "*", com.azure.core.util.Context.NONE);
    }
}
