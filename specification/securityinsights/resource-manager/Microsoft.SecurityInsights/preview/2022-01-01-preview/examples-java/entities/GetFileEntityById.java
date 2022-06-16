import com.azure.core.util.Context;

/** Samples for Entities Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/entities/GetFileEntityById.json
     */
    /**
     * Sample code: Get a file entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAFileEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "af378b21-b4aa-4fe7-bc70-13f8621a322f", Context.NONE);
    }
}
