
import com.azure.resourcemanager.storagemover.models.DayOfWeek;
import com.azure.resourcemanager.storagemover.models.Minute;
import com.azure.resourcemanager.storagemover.models.Time;
import com.azure.resourcemanager.storagemover.models.UploadLimitSchedule;
import com.azure.resourcemanager.storagemover.models.UploadLimitWeeklyRecurrence;
import java.util.Arrays;

/**
 * Samples for Agents CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Agents_CreateOrUpdate_MaximumSet.json
     */
    /**
     * Sample code: Agents_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        agentsCreateOrUpdateMaximumSet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.agents().define("examples-agentName")
            .withExistingStorageMover("examples-rg", "examples-storageMoverName")
            .withArcResourceId(
                "/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName")
            .withArcVmUuid("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9").withDescription("Example Agent Description")
            .withUploadLimitSchedule(new UploadLimitSchedule().withWeeklyRecurrences(Arrays
                .asList(new UploadLimitWeeklyRecurrence().withStartTime(new Time().withHour(9).withMinute(Minute.ZERO))
                    .withEndTime(new Time().withHour(18).withMinute(Minute.THREE_ZERO))
                    .withDays(Arrays.asList(DayOfWeek.MONDAY)).withLimitInMbps(2000))))
            .create();
    }
}
