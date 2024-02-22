
/**
 * Samples for Workspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Workspaces_Delete.json
     */
    /**
     * Sample code: Workspaces_Delete.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void workspacesDelete(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.workspaces().deleteWithResponse("contoso-resources", "contoso", "default",
            com.azure.core.util.Context.NONE);
    }
}
