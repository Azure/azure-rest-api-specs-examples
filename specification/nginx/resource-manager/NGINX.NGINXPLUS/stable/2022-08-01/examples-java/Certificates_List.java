import com.azure.core.util.Context;

/** Samples for Certificates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Certificates_List.json
     */
    /**
     * Sample code: Certificates_List.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void certificatesList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.certificates().list("myResourceGroup", "myDeployment", Context.NONE);
    }
}
