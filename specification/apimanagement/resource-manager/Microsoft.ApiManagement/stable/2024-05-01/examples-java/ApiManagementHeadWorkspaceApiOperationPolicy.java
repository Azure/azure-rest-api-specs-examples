
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceApiOperationPolicy GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadWorkspaceApiOperationPolicy.json
     */
    /**
     * Sample code: ApiManagementHeadWorkspaceApiOperationPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadWorkspaceApiOperationPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiOperationPolicies().getEntityTagWithResponse("rg1", "apimService1", "wks1",
            "5600b539c53f5b0062040001", "5600b53ac53f5b0062080006", PolicyIdName.POLICY,
            com.azure.core.util.Context.NONE);
    }
}
