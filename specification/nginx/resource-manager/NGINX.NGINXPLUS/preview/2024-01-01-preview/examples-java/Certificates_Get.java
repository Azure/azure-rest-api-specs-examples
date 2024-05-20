
/**
 * Samples for Certificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Certificates_Get.json
     */
    /**
     * Sample code: Certificates_Get.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesGet(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().getWithResponse("myResourceGroup", "myDeployment", "default",
            com.azure.core.util.Context.NONE);
    }
}
