
/**
 * Samples for ApiVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/
     * ApiVersions_CreateOrUpdate.json
     */
    /**
     * Sample code: ApiVersions_CreateOrUpdate.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiVersionsCreateOrUpdate(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiVersions().define("2023-01-01")
            .withExistingApi("contoso-resources", "contoso", "default", "echo-api").create();
    }
}
