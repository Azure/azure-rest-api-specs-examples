import com.azure.core.util.Context;

/** Samples for Entities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/GetDnsEntityById.json
     */
    /**
     * Sample code: Get a dns entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getADnsEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "f4e74920-f2c0-4412-a45f-66d94fdf01f8", Context.NONE);
    }
}
