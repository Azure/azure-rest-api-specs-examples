
/**
 * Samples for Apis Head.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Apis_Head.json
     */
    /**
     * Sample code: Apis_Head.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apisHead(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apis().headWithResponse("contoso-resources", "contoso", "default", "echo-api",
            com.azure.core.util.Context.NONE);
    }
}
