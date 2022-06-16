import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.Transformation;

/** Samples for Transformations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Transformation_Update.json
     */
    /**
     * Sample code: Update a transformation.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateATransformation(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Transformation resource =
            manager
                .transformations()
                .getWithResponse("sjrg4423", "sj8374", "transformation952", Context.NONE)
                .getValue();
        resource.update().withQuery("New query").apply();
    }
}
