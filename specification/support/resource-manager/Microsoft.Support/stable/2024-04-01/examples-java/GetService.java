
/**
 * Samples for Services Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/GetService.json
     */
    /**
     * Sample code: Gets details of the Azure service.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getsDetailsOfTheAzureService(com.azure.resourcemanager.support.SupportManager manager) {
        manager.services().getWithResponse("service_guid", com.azure.core.util.Context.NONE);
    }
}
