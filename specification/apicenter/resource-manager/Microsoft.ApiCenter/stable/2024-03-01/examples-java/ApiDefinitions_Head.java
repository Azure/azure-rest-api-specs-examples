
/**
 * Samples for ApiDefinitions Head.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiDefinitions_Head.json
     */
    /**
     * Sample code: ApiDefinitions_Head.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiDefinitionsHead(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiDefinitions().headWithResponse("contoso-resources", "contoso", "default", "echo-api", "2023-01-01",
            "openapi", com.azure.core.util.Context.NONE);
    }
}
