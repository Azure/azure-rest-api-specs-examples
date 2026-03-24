
/**
 * Samples for Certificates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/Certificates_List.json
     */
    /**
     * Sample code: Certificates_List.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().list("myResourceGroup", "myDeployment", com.azure.core.util.Context.NONE);
    }
}
