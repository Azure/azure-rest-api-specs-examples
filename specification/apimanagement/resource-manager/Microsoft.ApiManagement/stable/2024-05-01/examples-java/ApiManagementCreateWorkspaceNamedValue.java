
import com.azure.resourcemanager.apimanagement.models.NamedValueCreateContract;
import java.util.Arrays;

/**
 * Samples for WorkspaceNamedValue CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceNamedValue.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceNamedValue.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceNamedValue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().createOrUpdate(
            "rg1", "apimService1", "wks1", "testprop2", new NamedValueCreateContract().withDisplayName("prop3name")
                .withValue("propValue").withTags(Arrays.asList("foo", "bar")).withSecret(false),
            null, com.azure.core.util.Context.NONE);
    }
}
