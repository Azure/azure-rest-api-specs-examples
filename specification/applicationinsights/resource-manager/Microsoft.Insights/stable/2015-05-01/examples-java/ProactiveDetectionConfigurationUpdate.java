import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.fluent.models.ApplicationInsightsComponentProactiveDetectionConfigurationInner;
import com.azure.resourcemanager.applicationinsights.models.ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions;
import java.util.Arrays;

/** Samples for ProactiveDetectionConfigurations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationUpdate.json
     */
    /**
     * Sample code: ProactiveDetectionConfigurationUpdate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void proactiveDetectionConfigurationUpdate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .proactiveDetectionConfigurations()
            .updateWithResponse(
                "my-resource-group",
                "my-component",
                "slowpageloadtime",
                new ApplicationInsightsComponentProactiveDetectionConfigurationInner()
                    .withName("slowpageloadtime")
                    .withEnabled(true)
                    .withSendEmailsToSubscriptionOwners(true)
                    .withCustomEmails(Arrays.asList("foo@microsoft.com", "foo2@microsoft.com"))
                    .withRuleDefinitions(
                        new ApplicationInsightsComponentProactiveDetectionConfigurationRuleDefinitions()
                            .withName("slowpageloadtime")
                            .withDisplayName("Slow page load time")
                            .withDescription("Smart Detection rules notify you of performance anomaly issues.")
                            .withHelpUrl(
                                "https://docs.microsoft.com/en-us/azure/application-insights/app-insights-proactive-performance-diagnostics")
                            .withIsHidden(false)
                            .withIsEnabledByDefault(true)
                            .withIsInPreview(false)
                            .withSupportsEmailNotifications(true)),
                Context.NONE);
    }
}
