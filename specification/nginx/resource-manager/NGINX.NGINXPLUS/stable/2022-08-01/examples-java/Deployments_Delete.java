import com.azure.core.util.Context;

/** Samples for Deployments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Deployments_Delete.json
     */
    /**
     * Sample code: Deployments_Delete.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsDelete(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().delete("myResourceGroup", "myDeployment", Context.NONE);
    }
}
