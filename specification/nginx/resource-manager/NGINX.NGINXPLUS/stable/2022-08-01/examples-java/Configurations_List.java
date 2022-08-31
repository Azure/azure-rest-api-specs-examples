import com.azure.core.util.Context;

/** Samples for Configurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Configurations_List.json
     */
    /**
     * Sample code: Configurations_List.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().list("myResourceGroup", "myDeployment", Context.NONE);
    }
}
