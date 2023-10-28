import com.azure.resourcemanager.applicationinsights.models.ComponentPurgeBody;
import com.azure.resourcemanager.applicationinsights.models.ComponentPurgeBodyFilters;
import java.util.Arrays;

/** Samples for Components Purge. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsPurge.json
     */
    /**
     * Sample code: ComponentPurge.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentPurge(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .components()
            .purgeWithResponse(
                "OIAutoRest5123",
                "aztest5048",
                new ComponentPurgeBody()
                    .withTable("Heartbeat")
                    .withFilters(
                        Arrays
                            .asList(
                                new ComponentPurgeBodyFilters()
                                    .withColumn("TimeGenerated")
                                    .withOperator(">")
                                    .withValue("2017-09-01T00:00:00"))),
                com.azure.core.util.Context.NONE);
    }
}
