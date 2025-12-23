
import java.util.Arrays;

/**
 * Samples for DataExports CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/DataExportCreateOrUpdate.json
     */
    /**
     * Sample code: DataExportCreate.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void dataExportCreate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.dataExports().define("export1").withExistingWorkspace("RgTest1", "DeWnTest1234")
            .withTableNames(Arrays.asList("Heartbeat"))
            .withResourceId(
                "/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test")
            .create();
    }
}
