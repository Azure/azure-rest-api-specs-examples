import com.azure.core.util.Context;

/** Samples for Certificates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Certificates_Delete.json
     */
    /**
     * Sample code: Certificates_Delete.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesDelete(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().delete("myResourceGroup", "myDeployment", "default", Context.NONE);
    }
}
