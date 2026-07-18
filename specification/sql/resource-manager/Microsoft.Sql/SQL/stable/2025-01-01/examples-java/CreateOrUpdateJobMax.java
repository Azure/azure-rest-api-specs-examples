
import com.azure.resourcemanager.sql.fluent.models.JobInner;
import com.azure.resourcemanager.sql.models.JobSchedule;
import com.azure.resourcemanager.sql.models.JobScheduleType;
import java.time.OffsetDateTime;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateJobMax.json
     */
    /**
     * Sample code: Create a job with all properties specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createAJobWithAllPropertiesSpecified(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobs().createOrUpdateWithResponse("group1", "server1", "agent1", "job1",
            new JobInner().withDescription("my favourite job")
                .withSchedule(new JobSchedule().withStartTime(OffsetDateTime.parse("2015-09-24T18:30:01Z"))
                    .withEndTime(OffsetDateTime.parse("2015-09-24T23:59:59Z")).withType(JobScheduleType.RECURRING)
                    .withEnabled(true).withInterval("PT5M")),
            com.azure.core.util.Context.NONE);
    }
}
