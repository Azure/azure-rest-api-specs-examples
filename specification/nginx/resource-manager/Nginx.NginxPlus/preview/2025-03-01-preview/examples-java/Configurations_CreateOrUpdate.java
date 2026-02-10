
/**
 * Samples for Configurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Configurations_CreateOrUpdate.json
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
