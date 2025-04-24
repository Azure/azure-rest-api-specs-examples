
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceApiPolicy GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceApiPolicy.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceApiPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspaceApiPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiPolicies().getEntityTagWithResponse("rg1", "apimService1", "wks1",
            "57d1f7558aa04f15146d9d8a", PolicyIdName.POLICY, com.azure.core.util.Context.NONE);
    }
}
