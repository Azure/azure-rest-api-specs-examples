
import com.azure.resourcemanager.apimanagement.fluent.models.ApiReleaseContractInner;

/**
 * Samples for WorkspaceApiRelease CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceApiRelease.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceApiRelease.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceApiRelease(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiReleases().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "a1", "testrev",
            new ApiReleaseContractInner().withApiId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/a1")
                .withNotes("yahooagain"),
            null, com.azure.core.util.Context.NONE);
    }
}
