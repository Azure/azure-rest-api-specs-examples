
/**
 * Samples for Environments Head.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Environments_Head.json
     */
    /**
     * Sample code: Environments_Head.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void environmentsHead(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.environments().headWithResponse("contoso-resources", "contoso", "default", "public",
            com.azure.core.util.Context.NONE);
    }
}
