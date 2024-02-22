
/**
 * Samples for ApiVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiVersions_List.json
     */
    /**
     * Sample code: ApiVersions_ListByApi.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiVersionsListByApi(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiVersions().list("contoso-resources", "contoso", "default", "echo-api", null,
            com.azure.core.util.Context.NONE);
    }
}
