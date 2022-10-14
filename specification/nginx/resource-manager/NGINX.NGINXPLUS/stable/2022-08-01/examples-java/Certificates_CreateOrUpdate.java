/** Samples for Certificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Certificates_CreateOrUpdate.json
     */
    /**
     * Sample code: Certificates_CreateOrUpdate.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesCreateOrUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager
            .certificates()
            .define("default")
            .withRegion((String) null)
            .withExistingNginxDeployment("myResourceGroup", "myDeployment")
            .create();
    }
}
