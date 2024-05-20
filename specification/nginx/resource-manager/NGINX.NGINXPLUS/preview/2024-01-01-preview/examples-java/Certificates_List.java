
/**
 * Samples for Certificates List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Certificates_List.json
     */
    /**
     * Sample code: Certificates_List.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().list("myResourceGroup", "myDeployment", com.azure.core.util.Context.NONE);
    }
}
