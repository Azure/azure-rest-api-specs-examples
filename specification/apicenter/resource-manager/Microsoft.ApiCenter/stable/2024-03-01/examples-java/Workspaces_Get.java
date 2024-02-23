
/**
 * Samples for Workspaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Workspaces_Get.json
     */
    /**
     * Sample code: Workspaces_Get.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void workspacesGet(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.workspaces().getWithResponse("contoso-resources", "contoso", "default",
            com.azure.core.util.Context.NONE);
    }
}
