import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.fluent.models.OriginGroupInner;
import com.azure.resourcemanager.cdn.models.HealthProbeParameters;
import com.azure.resourcemanager.cdn.models.HealthProbeRequestType;
import com.azure.resourcemanager.cdn.models.ProbeProtocol;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import com.azure.resourcemanager.cdn.models.ResponseBasedDetectedErrorTypes;
import com.azure.resourcemanager.cdn.models.ResponseBasedOriginErrorDetectionParameters;
import java.util.Arrays;

/** Samples for OriginGroups Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/OriginGroups_Create.json
     */
    /**
     * Sample code: OriginGroups_Create.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originGroupsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getOriginGroups()
            .create(
                "RG",
                "profile1",
                "endpoint1",
                "origingroup1",
                new OriginGroupInner()
                    .withHealthProbeSettings(
                        new HealthProbeParameters()
                            .withProbePath("/health.aspx")
                            .withProbeRequestType(HealthProbeRequestType.GET)
                            .withProbeProtocol(ProbeProtocol.HTTP)
                            .withProbeIntervalInSeconds(120))
                    .withOrigins(
                        Arrays
                            .asList(
                                new ResourceReference()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin1")))
                    .withResponseBasedOriginErrorDetectionSettings(
                        new ResponseBasedOriginErrorDetectionParameters()
                            .withResponseBasedDetectedErrorTypes(ResponseBasedDetectedErrorTypes.TCP_ERRORS_ONLY)
                            .withResponseBasedFailoverThresholdPercentage(10)),
                Context.NONE);
    }
}
