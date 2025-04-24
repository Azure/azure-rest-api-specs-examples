
import com.azure.resourcemanager.apimanagement.fluent.models.ApiVersionSetContractInner;
import com.azure.resourcemanager.apimanagement.models.VersioningScheme;

/**
 * Samples for WorkspaceApiVersionSet CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceApiVersionSet.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceApiVersionSet.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspaceApiVersionSet(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiVersionSets().createOrUpdateWithResponse(
            "rg1", "apimService1", "wks1", "api1", new ApiVersionSetContractInner().withDisplayName("api set 1")
                .withVersioningScheme(VersioningScheme.SEGMENT).withDescription("Version configuration"),
            null, com.azure.core.util.Context.NONE);
    }
}
