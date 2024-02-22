
/**
 * Samples for Deployments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Deployments_List.json
     */
    /**
     * Sample code: Deployments_ListByApi.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void deploymentsListByApi(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.deployments().list("contoso-resources", "contoso", "default", "echo-api", null,
            com.azure.core.util.Context.NONE);
    }
}
