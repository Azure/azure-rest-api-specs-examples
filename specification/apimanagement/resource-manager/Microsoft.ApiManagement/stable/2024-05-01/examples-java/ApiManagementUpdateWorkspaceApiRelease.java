
import com.azure.resourcemanager.apimanagement.fluent.models.ApiReleaseContractInner;

/**
 * Samples for WorkspaceApiRelease Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspaceApiRelease.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspaceApiRelease.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspaceApiRelease(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiReleases().updateWithResponse("rg1", "apimService1", "wks1", "a1", "testrev", "*",
            new ApiReleaseContractInner().withApiId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/a1")
                .withNotes("yahooagain"),
            com.azure.core.util.Context.NONE);
    }
}
