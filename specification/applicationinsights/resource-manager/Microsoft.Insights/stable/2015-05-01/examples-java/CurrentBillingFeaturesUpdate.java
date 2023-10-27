import com.azure.resourcemanager.applicationinsights.fluent.models.ApplicationInsightsComponentBillingFeaturesInner;
import com.azure.resourcemanager.applicationinsights.models.ApplicationInsightsComponentDataVolumeCap;
import java.util.Arrays;

/** Samples for ComponentCurrentBillingFeatures Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/CurrentBillingFeaturesUpdate.json
     */
    /**
     * Sample code: ComponentCurrentBillingFeaturesUpdate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentCurrentBillingFeaturesUpdate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .componentCurrentBillingFeatures()
            .updateWithResponse(
                "my-resource-group",
                "my-component",
                new ApplicationInsightsComponentBillingFeaturesInner()
                    .withDataVolumeCap(
                        new ApplicationInsightsComponentDataVolumeCap()
                            .withCap(100.0F)
                            .withStopSendNotificationWhenHitCap(true))
                    .withCurrentBillingFeatures(Arrays.asList("Basic", "Application Insights Enterprise")),
                com.azure.core.util.Context.NONE);
    }
}
