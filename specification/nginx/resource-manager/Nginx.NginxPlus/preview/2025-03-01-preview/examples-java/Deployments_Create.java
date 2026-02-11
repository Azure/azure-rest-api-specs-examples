
/**
 * Samples for Deployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Deployments_Create.json
     */
    /**
     * Sample code: Deployments_Create.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsCreate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().define("myDeployment").withRegion((String) null)
            .withExistingResourceGroup("myResourceGroup").create();
    }
}
