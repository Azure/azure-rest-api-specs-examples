
/**
 * Samples for Deployments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Deployments_List.json
     */
    /**
     * Sample code: Deployments_List.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().list(com.azure.core.util.Context.NONE);
    }
}
