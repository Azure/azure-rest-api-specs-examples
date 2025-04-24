
import com.azure.resourcemanager.apimanagement.models.TagCreateUpdateParameters;

/**
 * Samples for WorkspaceTag CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceTag.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceTag.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTags().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "tagId1",
            new TagCreateUpdateParameters().withDisplayName("tag1"), null, com.azure.core.util.Context.NONE);
    }
}
