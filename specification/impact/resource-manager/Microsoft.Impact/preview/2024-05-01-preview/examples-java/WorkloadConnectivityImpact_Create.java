
import com.azure.resourcemanager.impactreporting.models.ClientIncidentDetails;
import com.azure.resourcemanager.impactreporting.models.Connectivity;
import com.azure.resourcemanager.impactreporting.models.IncidentSource;
import com.azure.resourcemanager.impactreporting.models.Protocol;
import com.azure.resourcemanager.impactreporting.models.SourceOrTarget;
import com.azure.resourcemanager.impactreporting.models.Toolset;
import com.azure.resourcemanager.impactreporting.models.Workload;
import com.azure.resourcemanager.impactreporting.models.WorkloadImpactProperties;
import java.time.OffsetDateTime;

/**
 * Samples for WorkloadImpacts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/WorkloadConnectivityImpact_Create.json
     */
    /**
     * Sample code: Reporting a connectivity impact.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void
        reportingAConnectivityImpact(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.workloadImpacts().define("impact-001").withProperties(new WorkloadImpactProperties()
            .withStartDateTime(OffsetDateTime.parse("2022-06-15T05:59:46.6517821Z"))
            .withImpactedResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservercontext")
            .withImpactCategory("Resource.Connectivity").withImpactDescription("conection failure")
            .withConnectivity(new Connectivity().withProtocol(Protocol.TCP).withPort(1443)
                .withSource(new SourceOrTarget().withAzureResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceSub/providers/Microsoft.compute/virtualmachines/vm1"))
                .withTarget(new SourceOrTarget().withAzureResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceSub/providers/Microsoft.compute/virtualmachines/vm2")))
            .withWorkload(new Workload().withContext("webapp/scenario1").withToolset(Toolset.OTHER))
            .withClientIncidentDetails(new ClientIncidentDetails().withClientIncidentId("AA123")
                .withClientIncidentSource(IncidentSource.JIRA)))
            .create();
    }
}
