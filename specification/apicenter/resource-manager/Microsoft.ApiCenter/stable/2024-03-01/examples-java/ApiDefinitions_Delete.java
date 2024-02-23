
/**
 * Samples for ApiDefinitions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiDefinitions_Delete.
     * json
     */
    /**
     * Sample code: ApiDefinitions_Delete.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiDefinitionsDelete(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiDefinitions().deleteWithResponse("contoso-resources", "contoso", "default", "echo-api", "2023-01-01",
            "openapi", com.azure.core.util.Context.NONE);
    }
}
