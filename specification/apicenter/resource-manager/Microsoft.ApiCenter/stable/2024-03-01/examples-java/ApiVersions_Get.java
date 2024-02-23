
/**
 * Samples for ApiVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiVersions_Get.json
     */
    /**
     * Sample code: ApiVersions_Get.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiVersionsGet(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiVersions().getWithResponse("contoso-resources", "contoso", "default", "echo-api", "2023-01-01",
            com.azure.core.util.Context.NONE);
    }
}
