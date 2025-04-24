
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceProductPolicy Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceProductPolicy.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceProductPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceProductPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductPolicies().getWithResponse("rg1", "apimService1", "wks1", "kjoshiarmTemplateProduct4",
            PolicyIdName.POLICY, null, com.azure.core.util.Context.NONE);
    }
}
