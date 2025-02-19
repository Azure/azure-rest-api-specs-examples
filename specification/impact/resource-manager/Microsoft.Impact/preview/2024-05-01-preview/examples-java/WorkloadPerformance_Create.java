
import com.azure.resourcemanager.impactreporting.models.ClientIncidentDetails;
import com.azure.resourcemanager.impactreporting.models.IncidentSource;
import com.azure.resourcemanager.impactreporting.models.MetricUnit;
import com.azure.resourcemanager.impactreporting.models.Performance;
import com.azure.resourcemanager.impactreporting.models.Toolset;
import com.azure.resourcemanager.impactreporting.models.Workload;
import com.azure.resourcemanager.impactreporting.models.WorkloadImpactProperties;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for WorkloadImpacts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/WorkloadPerformance_Create.json
     */
    /**
     * Sample code: Reporting performance related impact.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void
        reportingPerformanceRelatedImpact(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.workloadImpacts().define("impact-002").withProperties(new WorkloadImpactProperties()
            .withStartDateTime(OffsetDateTime.parse("2022-06-15T05:59:46.6517821Z"))
            .withImpactedResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservercontext")
            .withImpactCategory("Resource.Performance").withImpactDescription("high cpu utilization")
            .withPerformance(Arrays.asList(new Performance().withMetricName("CPU").withExpected(60.0D).withActual(90.0D)
                .withUnit(MetricUnit.fromString("garbage"))))
            .withWorkload(new Workload().withContext("webapp/scenario1").withToolset(Toolset.OTHER))
            .withClientIncidentDetails(new ClientIncidentDetails().withClientIncidentId("AA123")
                .withClientIncidentSource(IncidentSource.JIRA)))
            .create();
    }
}
