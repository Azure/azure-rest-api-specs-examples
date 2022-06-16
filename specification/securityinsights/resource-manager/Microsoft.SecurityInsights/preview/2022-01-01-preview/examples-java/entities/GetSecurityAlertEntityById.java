import com.azure.core.util.Context;

/** Samples for Entities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/entities/GetSecurityAlertEntityById.json
     */
    /**
     * Sample code: Get a security alert entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getASecurityAlertEntity(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "4aa486e0-6f85-41af-99ea-7acdce7be6c8", Context.NONE);
    }
}
