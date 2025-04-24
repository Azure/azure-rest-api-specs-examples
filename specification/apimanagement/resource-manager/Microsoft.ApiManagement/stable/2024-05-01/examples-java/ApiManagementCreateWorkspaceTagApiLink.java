
import com.azure.resourcemanager.apimanagement.fluent.models.TagApiLinkContractInner;

/**
 * Samples for WorkspaceTagApiLink CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceTagApiLink.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceTagApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceTagApiLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagApiLinks().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "tag1", "link1",
            new TagApiLinkContractInner().withApiId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/echo-api"),
            com.azure.core.util.Context.NONE);
    }
}
