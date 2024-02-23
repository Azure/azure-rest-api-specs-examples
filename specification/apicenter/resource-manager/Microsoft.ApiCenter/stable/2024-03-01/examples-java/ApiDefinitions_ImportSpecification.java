
/**
 * Samples for ApiDefinitions ImportSpecification.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/
     * ApiDefinitions_ImportSpecification.json
     */
    /**
     * Sample code: ApiDefinitions_ImportSpecification.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiDefinitionsImportSpecification(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiDefinitions().importSpecification("contoso-resources", "contoso", "default", "echo-api",
            "2023-01-01", "openapi", null, com.azure.core.util.Context.NONE);
    }
}
