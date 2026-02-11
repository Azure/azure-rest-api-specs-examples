
import com.azure.resourcemanager.nginx.models.NginxDeployment;

/**
 * Samples for Deployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Deployments_UpdateSubnet.json
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
