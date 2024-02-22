
/**
 * Samples for Deployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/
     * Deployments_CreateOrUpdate.json
     */
    /**
     * Sample code: Deployments_CreateOrUpdate.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void deploymentsCreateOrUpdate(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.deployments().define("production")
            .withExistingApi("contoso-resources", "contoso", "default", "echo-api").create();
    }
}
