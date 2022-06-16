import com.azure.core.util.Context;

/** Samples for Entities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/entities/GetFileHashEntityById.json
     */
    /**
     * Sample code: Get a file hash entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAFileHashEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "ea359fa6-c1e5-f878-e105-6344f3e399a1", Context.NONE);
    }
}
