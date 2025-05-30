
import com.azure.resourcemanager.synapse.fluent.models.MaintenanceWindowsInner;
import com.azure.resourcemanager.synapse.models.DayOfWeek;
import com.azure.resourcemanager.synapse.models.MaintenanceWindowTimeRange;
import java.util.Arrays;

/**
 * Samples for SqlPoolMaintenanceWindows CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/
     * CreateOrUpdateMaintenanceWindows.json
     */
    /**
     * Sample code: Sets maintenance window settings for a selected SQL Analytics Pool.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void setsMaintenanceWindowSettingsForASelectedSQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolMaintenanceWindows().createOrUpdateWithResponse("samplerg", "testworkspace", "testsp", "current",
            new MaintenanceWindowsInner().withTimeRanges(Arrays.asList(new MaintenanceWindowTimeRange()
                .withDayOfWeek(DayOfWeek.SATURDAY).withStartTime("00:00:00").withDuration("PT60M"))),
            com.azure.core.util.Context.NONE);
    }
}
