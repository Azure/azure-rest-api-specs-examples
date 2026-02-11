
import com.azure.resourcemanager.nginx.models.NginxDeployment;

/**
 * Samples for Deployments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Deployments_Update.json
     */
    /**
     * Sample code: Deployments_Update.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        NginxDeployment resource = manager.deployments()
            .getByResourceGroupWithResponse("myResourceGroup", "myDeployment", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
