
/**
 * Samples for ApiDefinitions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/
     * ApiDefinitions_CreateOrUpdate.json
     */
    /**
     * Sample code: ApiDefinitions_CreateOrUpdate.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiDefinitionsCreateOrUpdate(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiDefinitions().define("openapi")
            .withExistingVersion("contoso-resources", "contoso", "default", "openapi", "2023-01-01").create();
    }
}
