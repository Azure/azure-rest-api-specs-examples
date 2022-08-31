import com.azure.core.util.Context;

/** Samples for Deployments GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Deployments_Get.json
     */
    /**
     * Sample code: Deployments_Get.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsGet(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().getByResourceGroupWithResponse("myResourceGroup", "myDeployment", Context.NONE);
    }
}
