
import com.azure.resourcemanager.apimanagement.fluent.models.ProductApiLinkContractInner;

/**
 * Samples for WorkspaceProductApiLink CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceProductApiLink.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceProductApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspaceProductApiLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductApiLinks().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "testproduct",
            "link1",
            new ProductApiLinkContractInner().withApiId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/echo-api"),
            com.azure.core.util.Context.NONE);
    }
}
