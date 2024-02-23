
/**
 * Samples for Services List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/
     * Services_ListBySubscription.json
     */
    /**
     * Sample code: Services_ListBySubscription.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void servicesListBySubscription(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.services().list(com.azure.core.util.Context.NONE);
    }
}
