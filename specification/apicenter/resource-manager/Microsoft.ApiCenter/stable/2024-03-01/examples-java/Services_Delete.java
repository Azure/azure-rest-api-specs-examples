
/**
 * Samples for Services Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Services_Delete.json
     */
    /**
     * Sample code: Services_Delete.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void servicesDelete(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.services().deleteByResourceGroupWithResponse("contoso-resources", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
