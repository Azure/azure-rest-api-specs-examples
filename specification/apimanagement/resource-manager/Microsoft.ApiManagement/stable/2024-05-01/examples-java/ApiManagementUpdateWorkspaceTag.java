
import com.azure.resourcemanager.apimanagement.models.TagCreateUpdateParameters;

/**
 * Samples for WorkspaceTag Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspaceTag.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspaceTag.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspaceTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceTags().updateWithResponse("rg1", "apimService1", "wks1", "temptag", "*",
            new TagCreateUpdateParameters().withDisplayName("temp tag"), com.azure.core.util.Context.NONE);
    }
}
