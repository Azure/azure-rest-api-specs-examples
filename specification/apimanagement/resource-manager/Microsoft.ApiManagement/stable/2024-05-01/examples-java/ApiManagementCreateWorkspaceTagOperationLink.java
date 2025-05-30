
import com.azure.resourcemanager.apimanagement.fluent.models.TagOperationLinkContractInner;

/**
 * Samples for WorkspaceTagOperationLink CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceTagOperationLink.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceTagOperationLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspaceTagOperationLink(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTagOperationLinks().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "tag1", "link1",
            new TagOperationLinkContractInner().withOperationId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/echo-api/operations/op1"),
            com.azure.core.util.Context.NONE);
    }
}
