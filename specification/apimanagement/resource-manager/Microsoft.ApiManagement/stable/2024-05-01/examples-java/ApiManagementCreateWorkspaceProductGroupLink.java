
import com.azure.resourcemanager.apimanagement.fluent.models.ProductGroupLinkContractInner;

/**
 * Samples for WorkspaceProductGroupLink CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceProductGroupLink.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceProductGroupLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspaceProductGroupLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceProductGroupLinks().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "testproduct",
            "link1",
            new ProductGroupLinkContractInner().withGroupId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/groups/group1"),
            com.azure.core.util.Context.NONE);
    }
}
