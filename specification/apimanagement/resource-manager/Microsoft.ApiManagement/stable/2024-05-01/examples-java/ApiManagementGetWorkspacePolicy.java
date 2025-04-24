
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspacePolicy Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspacePolicy.json
     */
    /**
     * Sample code: ApiManagementGetWorkspacePolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetWorkspacePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicies().getWithResponse("rg1", "apimService1", "wks1", PolicyIdName.POLICY, null,
            com.azure.core.util.Context.NONE);
    }
}
