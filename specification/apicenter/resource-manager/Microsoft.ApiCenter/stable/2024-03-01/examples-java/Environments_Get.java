
/**
 * Samples for Environments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Environments_Get.json
     */
    /**
     * Sample code: Environments_Get.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void environmentsGet(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.environments().getWithResponse("contoso-resources", "contoso", "default", "public",
            com.azure.core.util.Context.NONE);
    }
}
