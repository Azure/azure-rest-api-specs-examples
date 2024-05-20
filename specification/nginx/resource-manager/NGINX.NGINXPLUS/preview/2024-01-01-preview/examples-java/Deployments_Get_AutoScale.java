
/**
 * Samples for Deployments GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/
     * Deployments_Get_AutoScale.json
     */
    /**
     * Sample code: Deployments_Get_AutoScale.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsGetAutoScale(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().getByResourceGroupWithResponse("myResourceGroup", "myDeployment",
            com.azure.core.util.Context.NONE);
    }
}
