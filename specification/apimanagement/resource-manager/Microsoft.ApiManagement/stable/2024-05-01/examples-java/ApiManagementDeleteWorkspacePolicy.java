
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspacePolicy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspacePolicy.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspacePolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteWorkspacePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicies().deleteWithResponse("rg1", "apimService1", "wks1", PolicyIdName.POLICY, "*",
            com.azure.core.util.Context.NONE);
    }
}
