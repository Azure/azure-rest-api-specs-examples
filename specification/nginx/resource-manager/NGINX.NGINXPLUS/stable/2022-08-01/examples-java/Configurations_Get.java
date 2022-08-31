import com.azure.core.util.Context;

/** Samples for Configurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Configurations_Get.json
     */
    /**
     * Sample code: Configurations_Get.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsGet(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().getWithResponse("myResourceGroup", "myDeployment", "default", Context.NONE);
    }
}
