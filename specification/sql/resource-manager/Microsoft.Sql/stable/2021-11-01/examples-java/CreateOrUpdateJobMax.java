
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.JobInner;
import com.azure.resourcemanager.sql.models.JobSchedule;
import com.azure.resourcemanager.sql.models.JobScheduleType;
import java.time.OffsetDateTime;

/** Samples for Jobs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateJobMax.json
     */
    /**
     * Sample code: Create a job with all properties specified.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAJobWithAllPropertiesSpecified(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobs().createOrUpdateWithResponse("group1", "server1", "agent1",
            "job1",
            new JobInner().withDescription("my favourite job")
                .withSchedule(new JobSchedule().withStartTime(OffsetDateTime.parse("2015-09-24T18:30:01Z"))
                    .withEndTime(OffsetDateTime.parse("2015-09-24T23:59:59Z")).withType(JobScheduleType.RECURRING)
                    .withEnabled(true).withInterval("PT5M")),
            Context.NONE);
    }
}
