
/**
 * Samples for Workspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Workspaces_List.json
     */
    /**
     * Sample code: Workspaces_ListByService.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void workspacesListByService(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.workspaces().list("contoso-resources", "contoso", null, com.azure.core.util.Context.NONE);
    }
}
