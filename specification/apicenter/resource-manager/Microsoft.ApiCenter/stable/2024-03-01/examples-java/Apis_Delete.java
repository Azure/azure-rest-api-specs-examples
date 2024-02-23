
/**
 * Samples for Apis Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Apis_Delete.json
     */
    /**
     * Sample code: Apis_Delete.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apisDelete(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apis().deleteWithResponse("contoso-resources", "contoso", "default", "echo-api",
            com.azure.core.util.Context.NONE);
    }
}
