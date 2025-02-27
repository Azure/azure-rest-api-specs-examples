
/**
 * Samples for Configurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/
     * Configurations_CreateOrUpdate.json
     */
    /**
     * Sample code: Configurations_CreateOrUpdate.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsCreateOrUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().define("default").withExistingNginxDeployment("myResourceGroup", "myDeployment")
            .create();
    }
}
