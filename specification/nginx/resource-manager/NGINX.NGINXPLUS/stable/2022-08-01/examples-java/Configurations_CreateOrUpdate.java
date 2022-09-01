/** Samples for Configurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Configurations_CreateOrUpdate.json
     */
    /**
     * Sample code: Configurations_CreateOrUpdate.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsCreateOrUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager
            .configurations()
            .define("default")
            .withRegion((String) null)
            .withExistingNginxDeployment("myResourceGroup", "myDeployment")
            .create();
    }
}
