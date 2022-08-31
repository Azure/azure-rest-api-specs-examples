import com.azure.core.util.Context;

/** Samples for Deployments ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Deployments_ListByResourceGroup.json
     */
    /**
     * Sample code: Deployments_ListByResourceGroup.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsListByResourceGroup(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
