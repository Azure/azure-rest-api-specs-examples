
/**
 * Samples for Services GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Services_Get.json
     */
    /**
     * Sample code: Services_Get.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void servicesGet(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.services().getByResourceGroupWithResponse("contoso-resources", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
