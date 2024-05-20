
/**
 * Samples for Configurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Configurations_Delete.
     * json
     */
    /**
     * Sample code: Configurations_Delete.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsDelete(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().delete("myResourceGroup", "myDeployment", "default", com.azure.core.util.Context.NONE);
    }
}
