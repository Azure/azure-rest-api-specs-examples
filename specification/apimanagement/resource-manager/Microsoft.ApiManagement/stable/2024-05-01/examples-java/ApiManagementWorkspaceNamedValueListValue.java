
/**
 * Samples for WorkspaceNamedValue ListValue.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementWorkspaceNamedValueListValue.json
     */
    /**
     * Sample code: ApiManagementWorkspaceNamedValueListValue.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementWorkspaceNamedValueListValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceNamedValues().listValueWithResponse("rg1", "apimService1", "wks1",
            "testarmTemplateproperties2", com.azure.core.util.Context.NONE);
    }
}
