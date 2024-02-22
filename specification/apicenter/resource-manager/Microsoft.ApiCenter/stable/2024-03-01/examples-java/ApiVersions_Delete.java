
/**
 * Samples for ApiVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiVersions_Delete.json
     */
    /**
     * Sample code: ApiVersions_Delete.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiVersionsDelete(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiVersions().deleteWithResponse("contoso-resources", "contoso", "default", "echo-api", "2023-01-01",
            com.azure.core.util.Context.NONE);
    }
}
