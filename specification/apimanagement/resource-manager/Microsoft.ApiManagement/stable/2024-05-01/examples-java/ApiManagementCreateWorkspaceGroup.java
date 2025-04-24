
import com.azure.resourcemanager.apimanagement.models.GroupCreateParameters;

/**
 * Samples for WorkspaceGroup CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceGroup.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceGroup.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceGroup(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceGroups().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "tempgroup",
            new GroupCreateParameters().withDisplayName("temp group"), null, com.azure.core.util.Context.NONE);
    }
}
