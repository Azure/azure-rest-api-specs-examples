import com.azure.resourcemanager.deploymentmanager.models.ApiKeyAuthentication;
import com.azure.resourcemanager.deploymentmanager.models.HealthCheckStepProperties;
import com.azure.resourcemanager.deploymentmanager.models.RestAuthLocation;
import com.azure.resourcemanager.deploymentmanager.models.RestHealthCheck;
import com.azure.resourcemanager.deploymentmanager.models.RestHealthCheckStepAttributes;
import com.azure.resourcemanager.deploymentmanager.models.RestMatchQuantifier;
import com.azure.resourcemanager.deploymentmanager.models.RestRequest;
import com.azure.resourcemanager.deploymentmanager.models.RestRequestMethod;
import com.azure.resourcemanager.deploymentmanager.models.RestResponse;
import com.azure.resourcemanager.deploymentmanager.models.RestResponseRegex;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Steps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_health_check_createorupdate.json
     */
    /**
     * Sample code: Create health check step.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void createHealthCheckStep(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .steps()
            .define("healthCheckStep")
            .withRegion("centralus")
            .withExistingResourceGroup("myResourceGroup")
            .withProperties(
                new HealthCheckStepProperties()
                    .withAttributes(
                        new RestHealthCheckStepAttributes()
                            .withWaitDuration("PT15M")
                            .withMaxElasticDuration("PT30M")
                            .withHealthyStateDuration("PT2H")
                            .withHealthChecks(
                                Arrays
                                    .asList(
                                        new RestHealthCheck()
                                            .withName("appHealth")
                                            .withRequest(
                                                new RestRequest()
                                                    .withMethod(RestRequestMethod.GET)
                                                    .withUri(
                                                        "https://resthealth.healthservice.com/api/applications/contosoApp/healthStatus")
                                                    .withAuthentication(
                                                        new ApiKeyAuthentication()
                                                            .withName("Code")
                                                            .withIn(RestAuthLocation.QUERY)
                                                            .withValue(
                                                                "NBCapiMOBQyAAbCkeytoPadnvO0eGHmidwFz5rXpappznKp3Jt7LLg==")))
                                            .withResponse(
                                                new RestResponse()
                                                    .withSuccessStatusCodes(Arrays.asList("OK"))
                                                    .withRegex(
                                                        new RestResponseRegex()
                                                            .withMatches(
                                                                Arrays
                                                                    .asList(
                                                                        "(?i)Contoso-App",
                                                                        "(?i)\"health_status\":((.|\n"
                                                                            + ")*)\"(green|yellow)\"",
                                                                        "(?mi)^(\"application_host\": 94781052)$"))
                                                            .withMatchQuantifier(RestMatchQuantifier.ALL))),
                                        new RestHealthCheck()
                                            .withName("serviceHealth")
                                            .withRequest(
                                                new RestRequest()
                                                    .withMethod(RestRequestMethod.GET)
                                                    .withUri(
                                                        "https://resthealth.healthservice.com/api/services/contosoService/healthStatus")
                                                    .withAuthentication(
                                                        new ApiKeyAuthentication()
                                                            .withName("code")
                                                            .withIn(RestAuthLocation.HEADER)
                                                            .withValue(
                                                                "NBCapiMOBQyAAbCkeytoPadnvO0eGHmidwFz5rXpappznKp3Jt7LLg==")))
                                            .withResponse(
                                                new RestResponse()
                                                    .withSuccessStatusCodes(Arrays.asList("OK"))
                                                    .withRegex(
                                                        new RestResponseRegex()
                                                            .withMatches(
                                                                Arrays
                                                                    .asList(
                                                                        "(?i)Contoso-Service-EndToEnd",
                                                                        "(?i)\"health_status\":((.|\n)*)\"(green)\""))
                                                            .withMatchQuantifier(RestMatchQuantifier.ALL)))))))
            .withTags(mapOf())
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
