
import com.azure.resourcemanager.apimanagement.models.AlwaysLog;
import com.azure.resourcemanager.apimanagement.models.BodyDiagnosticSettings;
import com.azure.resourcemanager.apimanagement.models.DiagnosticUpdateContract;
import com.azure.resourcemanager.apimanagement.models.HttpMessageDiagnostic;
import com.azure.resourcemanager.apimanagement.models.PipelineDiagnosticSettings;
import com.azure.resourcemanager.apimanagement.models.SamplingSettings;
import com.azure.resourcemanager.apimanagement.models.SamplingType;
import java.util.Arrays;

/**
 * Samples for WorkspaceDiagnostic Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspaceDiagnostic.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspaceDiagnostic.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspaceDiagnostic(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceDiagnostics().updateWithResponse("rg1", "apimService1", "wks1", "applicationinsights", "*",
            new DiagnosticUpdateContract().withAlwaysLog(AlwaysLog.ALL_ERRORS)
                .withLoggerId("/workspaces/wks1/loggers/applicationinsights")
                .withSampling(new SamplingSettings().withSamplingType(SamplingType.FIXED).withPercentage(50.0D))
                .withFrontend(new PipelineDiagnosticSettings()
                    .withRequest(new HttpMessageDiagnostic().withHeaders(Arrays.asList("Content-type"))
                        .withBody(new BodyDiagnosticSettings().withBytes(512)))
                    .withResponse(new HttpMessageDiagnostic().withHeaders(Arrays.asList("Content-type"))
                        .withBody(new BodyDiagnosticSettings().withBytes(512))))
                .withBackend(new PipelineDiagnosticSettings()
                    .withRequest(new HttpMessageDiagnostic().withHeaders(Arrays.asList("Content-type"))
                        .withBody(new BodyDiagnosticSettings().withBytes(512)))
                    .withResponse(new HttpMessageDiagnostic().withHeaders(Arrays.asList("Content-type"))
                        .withBody(new BodyDiagnosticSettings().withBytes(512)))),
            com.azure.core.util.Context.NONE);
    }
}
