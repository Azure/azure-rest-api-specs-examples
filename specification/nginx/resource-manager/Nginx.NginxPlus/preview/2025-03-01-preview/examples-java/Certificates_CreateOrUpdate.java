
/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Certificates_CreateOrUpdate.json
     */
    /**
     * Sample code: Certificates_CreateOrUpdate.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesCreateOrUpdate(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().define("default").withExistingNginxDeployment("myResourceGroup", "myDeployment")
            .create();
    }
}
