
/**
 * Samples for Deployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Deployments_Delete.json
     */
    /**
     * Sample code: Deployments_Delete.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsDelete(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().delete("myResourceGroup", "myDeployment", com.azure.core.util.Context.NONE);
    }
}
