
/**
 * Samples for Certificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Certificates_Delete.json
     */
    /**
     * Sample code: Certificates_Delete.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesDelete(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().delete("myResourceGroup", "myDeployment", "default", com.azure.core.util.Context.NONE);
    }
}
