
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspacePolicy GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspacePolicy.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspacePolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadWorkspacePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicies().getEntityTagWithResponse("rg1", "apimService1", "wks1", PolicyIdName.POLICY,
            com.azure.core.util.Context.NONE);
    }
}
