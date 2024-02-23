
/**
 * Samples for Deployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Deployments_Get.json
     */
    /**
     * Sample code: Deployments_Get.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void deploymentsGet(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.deployments().getWithResponse("contoso-resources", "contoso", "default", "echo-api", "production",
            com.azure.core.util.Context.NONE);
    }
}
