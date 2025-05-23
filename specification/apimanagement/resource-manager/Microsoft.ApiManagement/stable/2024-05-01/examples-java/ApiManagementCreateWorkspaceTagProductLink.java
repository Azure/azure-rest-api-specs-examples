
import com.azure.resourcemanager.apimanagement.fluent.models.TagProductLinkContractInner;

/**
 * Samples for WorkspaceTagProductLink CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceTagProductLink.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceTagProductLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspaceTagProductLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagProductLinks().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "tag1", "link1",
            new TagProductLinkContractInner().withProductId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/products/product1"),
            com.azure.core.util.Context.NONE);
    }
}
