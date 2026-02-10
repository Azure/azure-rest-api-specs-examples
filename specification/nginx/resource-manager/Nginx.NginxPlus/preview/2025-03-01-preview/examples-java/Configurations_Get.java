
/**
 * Samples for Configurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Configurations_Get.json
     */
    /**
     * Sample code: Configurations_Get.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsGet(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().getWithResponse("myResourceGroup", "myDeployment", "default",
            com.azure.core.util.Context.NONE);
    }
}
