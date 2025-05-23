
import com.azure.resourcemanager.securityinsights.models.EntityGetInsightsParameters;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.UUID;

/**
 * Samples for Entities GetInsights.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * entities/insights/PostGetInsights.json
     */
    /**
     * Sample code: Entity Insight.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void entityInsight(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getInsightsWithResponse("myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1",
            new EntityGetInsightsParameters().withStartTime(OffsetDateTime.parse("2021-09-01T00:00:00.000Z"))
                .withEndTime(OffsetDateTime.parse("2021-10-01T00:00:00.000Z")).withAddDefaultExtendedTimeRange(false)
                .withInsightQueryIds(Arrays.asList(UUID.fromString("cae8d0aa-aa45-4d53-8d88-17dd64ffd4e4"))),
            com.azure.core.util.Context.NONE);
    }
}
