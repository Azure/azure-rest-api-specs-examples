
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceApiPolicy Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceApiPolicy.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceApiPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspaceApiPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiPolicies().getWithResponse("rg1", "apimService1", "wks1", "5600b59475ff190048040001",
            PolicyIdName.POLICY, null, com.azure.core.util.Context.NONE);
    }
}
