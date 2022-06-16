import com.azure.core.util.Context;

/** Samples for SqlPoolMaintenanceWindows Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetMaintenanceWindows.json
     */
    /**
     * Sample code: Gets maintenance window settings for a selected SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsMaintenanceWindowSettingsForASelectedSQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolMaintenanceWindows()
            .getWithResponse("samplerg", "testworkspace", "testsp", "current", Context.NONE);
    }
}
