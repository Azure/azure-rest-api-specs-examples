import com.azure.core.util.Context;

/** Samples for WorkspacePurge GetPurgeStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesPurgeOperation.json
     */
    /**
     * Sample code: WorkspacePurgeOperation.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacePurgeOperation(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .workspacePurges()
            .getPurgeStatusWithResponse(
                "OIAutoRest5123", "aztest5048", "purge-970318e7-b859-4edb-8903-83b1b54d0b74", Context.NONE);
    }
}
