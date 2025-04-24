
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceProductPolicy GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceProductPolicy.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceProductPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceProductPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductPolicies().getEntityTagWithResponse("rg1", "apimService1", "wks1", "unlimited",
            PolicyIdName.POLICY, com.azure.core.util.Context.NONE);
    }
}
