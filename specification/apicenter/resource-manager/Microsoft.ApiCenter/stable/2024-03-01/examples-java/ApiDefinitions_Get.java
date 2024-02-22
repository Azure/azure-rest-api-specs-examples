
/**
 * Samples for ApiDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiDefinitions_Get.json
     */
    /**
     * Sample code: ApiDefinitions_Get.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiDefinitionsGet(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiDefinitions().getWithResponse("contoso-resources", "contoso", "default", "echo-api", "2023-01-01",
            "openapi", com.azure.core.util.Context.NONE);
    }
}
