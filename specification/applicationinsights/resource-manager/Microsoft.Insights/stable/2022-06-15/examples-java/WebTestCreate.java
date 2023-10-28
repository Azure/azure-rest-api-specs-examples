import com.azure.resourcemanager.applicationinsights.models.WebTestGeolocation;
import com.azure.resourcemanager.applicationinsights.models.WebTestKind;
import com.azure.resourcemanager.applicationinsights.models.WebTestPropertiesConfiguration;
import java.util.Arrays;

/** Samples for WebTests CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestCreate.json
     */
    /**
     * Sample code: webTestCreate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestCreate(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .webTests()
            .define("my-webtest-my-component")
            .withRegion("South Central US")
            .withExistingResourceGroup("my-resource-group")
            .withKind(WebTestKind.PING)
            .withSyntheticMonitorId("my-webtest-my-component")
            .withWebTestName("my-webtest-my-component")
            .withDescription("Ping web test alert for mytestwebapp")
            .withEnabled(true)
            .withFrequency(900)
            .withTimeout(120)
            .withWebTestKind(WebTestKind.PING)
            .withRetryEnabled(true)
            .withLocations(Arrays.asList(new WebTestGeolocation().withLocation("us-fl-mia-edge")))
            .withConfiguration(
                new WebTestPropertiesConfiguration()
                    .withWebTest(
                        "<WebTest Name=\"my-webtest\" Id=\"678ddf96-1ab8-44c8-9274-123456789abc\" Enabled=\"True\""
                            + " CssProjectStructure=\"\" CssIteration=\"\" Timeout=\"120\" WorkItemIds=\"\""
                            + " xmlns=\"http://microsoft.com/schemas/VisualStudio/TeamTest/2010\" Description=\"\""
                            + " CredentialUserName=\"\" CredentialPassword=\"\" PreAuthenticate=\"True\""
                            + " Proxy=\"default\" StopOnError=\"False\" RecordedResultFile=\"\" ResultsLocale=\"\""
                            + " ><Items><Request Method=\"GET\" Guid=\"a4162485-9114-fcfc-e086-123456789abc\""
                            + " Version=\"1.1\" Url=\"http://my-component.azurewebsites.net\" ThinkTime=\"0\""
                            + " Timeout=\"120\" ParseDependentRequests=\"True\" FollowRedirects=\"True\""
                            + " RecordResult=\"True\" Cache=\"False\" ResponseTimeGoal=\"0\" Encoding=\"utf-8\""
                            + " ExpectedHttpStatusCode=\"200\" ExpectedResponseUrl=\"\" ReportingName=\"\""
                            + " IgnoreHttpStatusCode=\"False\" /></Items></WebTest>"))
            .create();
    }
}
