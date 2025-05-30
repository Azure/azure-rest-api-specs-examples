
import com.azure.resourcemanager.apimanagement.models.GroupUpdateParameters;

/**
 * Samples for WorkspaceGroup Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspaceGroup.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspaceGroup.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspaceGroup(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGroups().updateWithResponse("rg1", "apimService1", "wks1", "tempgroup", "*",
            new GroupUpdateParameters().withDisplayName("temp group"), com.azure.core.util.Context.NONE);
    }
}
