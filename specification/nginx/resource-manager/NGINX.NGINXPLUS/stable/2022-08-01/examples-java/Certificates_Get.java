import com.azure.core.util.Context;

/** Samples for Certificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Certificates_Get.json
     */
    /**
     * Sample code: Certificates_Get.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesGet(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().getWithResponse("myResourceGroup", "myDeployment", "default", Context.NONE);
    }
}
