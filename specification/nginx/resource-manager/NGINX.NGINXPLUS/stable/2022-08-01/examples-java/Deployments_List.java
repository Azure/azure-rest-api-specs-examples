import com.azure.core.util.Context;

/** Samples for Deployments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Deployments_List.json
     */
    /**
     * Sample code: Deployments_List.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().list(Context.NONE);
    }
}
