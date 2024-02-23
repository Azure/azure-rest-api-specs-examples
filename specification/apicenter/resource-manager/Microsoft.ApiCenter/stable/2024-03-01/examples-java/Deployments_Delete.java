
/**
 * Samples for Deployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Deployments_Delete.json
     */
    /**
     * Sample code: Deployments_Delete.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void deploymentsDelete(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.deployments().deleteWithResponse("contoso-resources", "contoso", "default", "echo-api", "production",
            com.azure.core.util.Context.NONE);
    }
}
