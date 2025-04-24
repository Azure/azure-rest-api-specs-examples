
import com.azure.resourcemanager.apimanagement.models.NamedValueUpdateParameters;
import java.util.Arrays;

/**
 * Samples for WorkspaceNamedValue Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspaceNamedValue.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspaceNamedValue.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspaceNamedValue(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().update("rg1", "apimService1", "wks1", "testprop2", "*",
            new NamedValueUpdateParameters().withDisplayName("prop3name").withValue("propValue")
                .withTags(Arrays.asList("foo", "bar2")).withSecret(false),
            com.azure.core.util.Context.NONE);
    }
}
