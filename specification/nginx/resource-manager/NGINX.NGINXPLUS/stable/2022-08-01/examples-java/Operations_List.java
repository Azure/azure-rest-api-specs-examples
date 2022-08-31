import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to NginxManager.
     */
    public static void operationsList(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.operations().list(Context.NONE);
    }
}
