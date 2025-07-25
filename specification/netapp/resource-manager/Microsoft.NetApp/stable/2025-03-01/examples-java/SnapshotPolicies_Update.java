
import com.azure.resourcemanager.netapp.models.DailySchedule;
import com.azure.resourcemanager.netapp.models.HourlySchedule;
import com.azure.resourcemanager.netapp.models.MonthlySchedule;
import com.azure.resourcemanager.netapp.models.SnapshotPolicy;
import com.azure.resourcemanager.netapp.models.WeeklySchedule;

/**
 * Samples for SnapshotPolicies Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/SnapshotPolicies_Update.json
     */
    /**
     * Sample code: SnapshotPolicies_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotPoliciesUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        SnapshotPolicy resource = manager.snapshotPolicies()
            .getWithResponse("myRG", "account1", "snapshotPolicyName", com.azure.core.util.Context.NONE).getValue();
        resource.update().withHourlySchedule(new HourlySchedule().withSnapshotsToKeep(2).withMinute(50))
            .withDailySchedule(new DailySchedule().withSnapshotsToKeep(4).withHour(14).withMinute(30))
            .withWeeklySchedule(
                new WeeklySchedule().withSnapshotsToKeep(3).withDay("Wednesday").withHour(14).withMinute(45))
            .withMonthlySchedule(
                new MonthlySchedule().withSnapshotsToKeep(5).withDaysOfMonth("10,11,12").withHour(14).withMinute(15))
            .withEnabled(true).apply();
    }
}
