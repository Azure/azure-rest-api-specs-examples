
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * entities/GetAzureResourceEntityById.json
     */
    /**
     * Sample code: Get an azure resource entity.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAnAzureResourceEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1",
            com.azure.core.util.Context.NONE);
    }
}
