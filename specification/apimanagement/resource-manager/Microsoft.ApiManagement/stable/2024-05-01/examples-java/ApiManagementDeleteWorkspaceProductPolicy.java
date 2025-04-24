
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceProductPolicy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceProductPolicy.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceProductPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceProductPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductPolicies().deleteWithResponse("rg1", "apimService1", "wks1", "testproduct",
            PolicyIdName.POLICY, "*", com.azure.core.util.Context.NONE);
    }
}
