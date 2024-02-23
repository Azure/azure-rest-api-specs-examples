
/**
 * Samples for Workspaces Head.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Workspaces_Head.json
     */
    /**
     * Sample code: Workspaces_Head.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void workspacesHead(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.workspaces().headWithResponse("contoso-resources", "contoso", "default",
            com.azure.core.util.Context.NONE);
    }
}
