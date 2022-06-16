import com.azure.core.util.Context;

/** Samples for Entities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entities/GetProcessEntityById.json
     */
    /**
     * Sample code: Get a process entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAProcessEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "7264685c-038c-42c6-948c-38e14ef1fb98", Context.NONE);
    }
}
