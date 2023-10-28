import com.azure.resourcemanager.applicationinsights.models.HeaderField;
import com.azure.resourcemanager.applicationinsights.models.WebTestGeolocation;
import com.azure.resourcemanager.applicationinsights.models.WebTestKind;
import com.azure.resourcemanager.applicationinsights.models.WebTestPropertiesRequest;
import com.azure.resourcemanager.applicationinsights.models.WebTestPropertiesValidationRules;
import java.util.Arrays;

/** Samples for WebTests CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestCreateStandard.json
     */
    /**
     * Sample code: webTestCreateStandard.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void webTestCreateStandard(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .webTests()
            .define("my-webtest-my-component")
            .withRegion("South Central US")
            .withExistingResourceGroup("my-resource-group")
            .withSyntheticMonitorId("my-webtest-my-component")
            .withWebTestName("my-webtest-my-component")
            .withDescription("Ping web test alert for mytestwebapp")
            .withEnabled(true)
            .withFrequency(900)
            .withTimeout(120)
            .withWebTestKind(WebTestKind.STANDARD)
            .withRetryEnabled(true)
            .withLocations(Arrays.asList(new WebTestGeolocation().withLocation("us-fl-mia-edge")))
            .withRequest(
                new WebTestPropertiesRequest()
                    .withRequestUrl("https://bing.com")
                    .withHeaders(
                        Arrays
                            .asList(
                                new HeaderField()
                                    .withHeaderFieldName("fakeTokenPlaceholder")
                                    .withHeaderFieldValue("de-DE"),
                                new HeaderField()
                                    .withHeaderFieldName("fakeTokenPlaceholder")
                                    .withHeaderFieldValue("de-DE")))
                    .withHttpVerb("POST")
                    .withRequestBody("SGVsbG8gd29ybGQ="))
            .withValidationRules(
                new WebTestPropertiesValidationRules().withSslCheck(true).withSslCertRemainingLifetimeCheck(100))
            .create();
    }
}
