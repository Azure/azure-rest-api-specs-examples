import com.azure.resourcemanager.redis.fluent.models.RedisPatchScheduleInner;
import com.azure.resourcemanager.redis.models.DayOfWeek;
import com.azure.resourcemanager.redis.models.DefaultName;
import com.azure.resourcemanager.redis.models.ScheduleEntry;
import java.time.Duration;
import java.util.Arrays;

/** Samples for PatchSchedules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCachePatchSchedulesCreateOrUpdate.json
     */
    /**
     * Sample code: RedisCachePatchSchedulesCreateOrUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCachePatchSchedulesCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .redisCaches()
            .manager()
            .serviceClient()
            .getPatchSchedules()
            .createOrUpdateWithResponse(
                "rg1",
                "cache1",
                DefaultName.DEFAULT,
                new RedisPatchScheduleInner()
                    .withScheduleEntries(
                        Arrays
                            .asList(
                                new ScheduleEntry()
                                    .withDayOfWeek(DayOfWeek.MONDAY)
                                    .withStartHourUtc(12)
                                    .withMaintenanceWindow(Duration.parse("PT5H")),
                                new ScheduleEntry().withDayOfWeek(DayOfWeek.TUESDAY).withStartHourUtc(12))),
                com.azure.core.util.Context.NONE);
    }
}
