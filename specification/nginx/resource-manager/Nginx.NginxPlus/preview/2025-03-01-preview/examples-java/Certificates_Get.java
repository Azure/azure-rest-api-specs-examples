
/**
 * Samples for Certificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Certificates_Get.json
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
