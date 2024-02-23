
/**
 * Samples for Environments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Environments_Delete.json
     */
    /**
     * Sample code: Environments_Delete.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void environmentsDelete(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.environments().deleteWithResponse("contoso-resources", "contoso", "default", "public",
            com.azure.core.util.Context.NONE);
    }
}
