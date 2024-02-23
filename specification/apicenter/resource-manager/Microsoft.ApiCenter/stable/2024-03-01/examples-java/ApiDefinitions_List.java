
/**
 * Samples for ApiDefinitions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiDefinitions_List.json
     */
    /**
     * Sample code: ApiDefinitions_ListByApiVersion.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apiDefinitionsListByApiVersion(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apiDefinitions().list("contoso-resources", "contoso", "default", "echo-api", "2023-01-01", null,
            com.azure.core.util.Context.NONE);
    }
}
