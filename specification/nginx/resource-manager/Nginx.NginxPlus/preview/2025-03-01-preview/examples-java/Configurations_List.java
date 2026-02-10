
/**
 * Samples for Configurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Configurations_List.json
     */
    /**
     * Sample code: Configurations_List.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().list("myResourceGroup", "myDeployment", com.azure.core.util.Context.NONE);
    }
}
