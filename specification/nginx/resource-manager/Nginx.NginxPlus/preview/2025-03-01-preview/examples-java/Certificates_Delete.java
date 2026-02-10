
/**
 * Samples for Certificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Certificates_Delete.json
     */
    /**
     * Sample code: Certificates_Delete.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesDelete(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().delete("myResourceGroup", "myDeployment", "default", com.azure.core.util.Context.NONE);
    }
}
