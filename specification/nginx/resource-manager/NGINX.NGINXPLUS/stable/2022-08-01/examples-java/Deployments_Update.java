import com.azure.core.util.Context;
import com.azure.resourcemanager.nginx.models.NginxDeployment;

/** Samples for Deployments Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Deployments_Update.json
     */
    /**
     * Sample code: Deployments_Update.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        NginxDeployment resource =
            manager
                .deployments()
                .getByResourceGroupWithResponse("myResourceGroup", "myDeployment", Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
