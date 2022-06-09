```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.EntityTimelineParameters;
import java.time.OffsetDateTime;

/** Samples for EntitiesGetTimeline List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entities/timeline/PostTimelineEntity.json
     */
    /**
     * Sample code: Entity timeline.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void entityTimeline(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entitiesGetTimelines()
            .listWithResponse(
                "myRg",
                "myWorkspace",
                "e1d3d618-e11f-478b-98e3-bb381539a8e1",
                new EntityTimelineParameters()
                    .withStartTime(OffsetDateTime.parse("2021-09-01T00:00:00.000Z"))
                    .withEndTime(OffsetDateTime.parse("2021-10-01T00:00:00.000Z"))
                    .withNumberOfBucket(4),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
