Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redis.fluent.models.RedisPatchScheduleInner;
import com.azure.resourcemanager.redis.models.DayOfWeek;
import com.azure.resourcemanager.redis.models.DefaultName;
import com.azure.resourcemanager.redis.models.ScheduleEntry;
import java.time.Duration;
import java.util.Arrays;

/** Samples for PatchSchedules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCachePatchSchedulesCreateOrUpdate.json
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
                Context.NONE);
    }
}
```
