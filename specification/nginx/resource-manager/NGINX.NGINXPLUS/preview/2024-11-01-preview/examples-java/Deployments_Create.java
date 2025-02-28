
/**
 * Samples for Deployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/Deployments_Create.json
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
