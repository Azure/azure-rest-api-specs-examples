
/**
 * Samples for ApiVersions Head.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiVersions_Head.json
     */
    /**
     * Sample code: ApiVersions_Head.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiVersionsHead(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiVersions().headWithResponse("contoso-resources", "contoso", "default", "echo-api", "2023-01-01",
            com.azure.core.util.Context.NONE);
    }
}
