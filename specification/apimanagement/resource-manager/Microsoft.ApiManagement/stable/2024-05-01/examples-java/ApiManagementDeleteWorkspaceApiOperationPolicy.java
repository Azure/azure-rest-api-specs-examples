
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceApiOperationPolicy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteWorkspaceApiOperationPolicy.json
     */
    /**
     * Sample code: ApiManagementDeleteWorkspaceApiOperationPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteWorkspaceApiOperationPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiOperationPolicies().deleteWithResponse("rg1", "apimService1", "wks1", "testapi",
            "testoperation", PolicyIdName.POLICY, "*", com.azure.core.util.Context.NONE);
    }
}
