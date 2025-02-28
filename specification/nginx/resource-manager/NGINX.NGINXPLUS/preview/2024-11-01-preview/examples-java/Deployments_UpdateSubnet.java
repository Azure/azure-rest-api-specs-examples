
import com.azure.resourcemanager.nginx.models.NginxDeployment;

/**
 * Samples for Deployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/Deployments_UpdateSubnet
     * .json
     */
    /**
     * Sample code: Deployments_UpdateSubnet.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsUpdateSubnet(com.azure.resourcemanager.nginx.NginxManager manager) {
        NginxDeployment resource = manager.deployments()
            .getByResourceGroupWithResponse("myResourceGroup", "myDeployment", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
