
/**
 * Samples for Services ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/
     * Services_ListByResourceGroup.json
     */
    /**
     * Sample code: Services_ListByResourceGroup.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void servicesListByResourceGroup(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.services().listByResourceGroup("contoso-resources", com.azure.core.util.Context.NONE);
    }
}
