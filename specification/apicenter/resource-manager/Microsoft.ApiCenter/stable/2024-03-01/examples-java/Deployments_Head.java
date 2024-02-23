
/**
 * Samples for Deployments Head.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Deployments_Head.json
     */
    /**
     * Sample code: Deployments_Head.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void deploymentsHead(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.deployments().headWithResponse("contoso-resources", "contoso", "default", "echo-api", "production",
            com.azure.core.util.Context.NONE);
    }
}
