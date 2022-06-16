import com.azure.core.util.Context;

/** Samples for Entities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entities/GetUrlEntityById.json
     */
    /**
     * Sample code: Get a url entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAUrlEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1", Context.NONE);
    }
}
