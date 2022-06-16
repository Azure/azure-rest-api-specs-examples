import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.EntityExpandParameters;
import java.time.OffsetDateTime;
import java.util.UUID;

/** Samples for Entities Expand. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/entities/expand/PostExpandEntity.json
     */
    /**
     * Sample code: Expand an entity.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void expandAnEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entities()
            .expandWithResponse(
                "myRg",
                "myWorkspace",
                "e1d3d618-e11f-478b-98e3-bb381539a8e1",
                new EntityExpandParameters()
                    .withEndTime(OffsetDateTime.parse("2019-05-26T00:00:00.000Z"))
                    .withExpansionId(UUID.fromString("a77992f3-25e9-4d01-99a4-5ff606cc410a"))
                    .withStartTime(OffsetDateTime.parse("2019-04-25T00:00:00.000Z")),
                Context.NONE);
    }
}
