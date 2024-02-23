
/**
 * Samples for Workspaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Workspaces_CreateOrUpdate
     * .json
     */
    /**
     * Sample code: Workspaces_CreateOrUpdate.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void workspacesCreateOrUpdate(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.workspaces().define("default").withExistingService("contoso-resources", "contoso").create();
    }
}
