import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.AlwaysLog;
import com.azure.resourcemanager.apimanagement.models.BodyDiagnosticSettings;
import com.azure.resourcemanager.apimanagement.models.DiagnosticContract;
import com.azure.resourcemanager.apimanagement.models.HttpMessageDiagnostic;
import com.azure.resourcemanager.apimanagement.models.PipelineDiagnosticSettings;
import com.azure.resourcemanager.apimanagement.models.SamplingSettings;
import com.azure.resourcemanager.apimanagement.models.SamplingType;
import java.util.Arrays;

/** Samples for ApiDiagnostic Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApiDiagnostic.json
     */
    /**
     * Sample code: ApiManagementUpdateApiDiagnostic.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApiDiagnostic(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        DiagnosticContract resource =
            manager
                .apiDiagnostics()
                .getWithResponse("rg1", "apimService1", "echo-api", "applicationinsights", Context.NONE)
                .getValue();
        resource
            .update()
            .withAlwaysLog(AlwaysLog.ALL_ERRORS)
            .withLoggerId("/loggers/applicationinsights")
            .withSampling(new SamplingSettings().withSamplingType(SamplingType.FIXED).withPercentage(50.0))
            .withFrontend(
                new PipelineDiagnosticSettings()
                    .withRequest(
                        new HttpMessageDiagnostic()
                            .withHeaders(Arrays.asList("Content-type"))
                            .withBody(new BodyDiagnosticSettings().withBytes(512)))
                    .withResponse(
                        new HttpMessageDiagnostic()
                            .withHeaders(Arrays.asList("Content-type"))
                            .withBody(new BodyDiagnosticSettings().withBytes(512))))
            .withBackend(
                new PipelineDiagnosticSettings()
                    .withRequest(
                        new HttpMessageDiagnostic()
                            .withHeaders(Arrays.asList("Content-type"))
                            .withBody(new BodyDiagnosticSettings().withBytes(512)))
                    .withResponse(
                        new HttpMessageDiagnostic()
                            .withHeaders(Arrays.asList("Content-type"))
                            .withBody(new BodyDiagnosticSettings().withBytes(512))))
            .withIfMatch("*")
            .apply();
    }
}
