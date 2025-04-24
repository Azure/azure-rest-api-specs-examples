
import com.azure.resourcemanager.apimanagement.models.PolicyFragmentContentFormat;

/**
 * Samples for WorkspacePolicyFragment Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspacePolicyFragmentFormat.json
     */
    /**
     * Sample code: ApiManagementGetWorkspacePolicyFragmentFormat.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetWorkspacePolicyFragmentFormat(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicyFragments().getWithResponse("rg1", "apimService1", "wks1", "policyFragment1",
            PolicyFragmentContentFormat.RAWXML, com.azure.core.util.Context.NONE);
    }
}
