
import com.azure.resourcemanager.cdn.models.HealthProbeParameters;
import com.azure.resourcemanager.cdn.models.HealthProbeRequestType;
import com.azure.resourcemanager.cdn.models.OriginGroupUpdateParameters;
import com.azure.resourcemanager.cdn.models.ProbeProtocol;
import com.azure.resourcemanager.cdn.models.ResourceReference;
import java.util.Arrays;

/** Samples for OriginGroups Update. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/OriginGroups_Update.json
     */
    /**
     * Sample code: OriginGroups_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originGroupsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getOriginGroups().update("RG", "profile1", "endpoint1",
            "originGroup1",
            new OriginGroupUpdateParameters()
                .withHealthProbeSettings(new HealthProbeParameters().withProbePath("/health.aspx")
                    .withProbeRequestType(HealthProbeRequestType.GET).withProbeProtocol(ProbeProtocol.HTTP)
                    .withProbeIntervalInSeconds(120))
                .withOrigins(Arrays.asList(new ResourceReference().withId(
                    "/subscriptions/subid/resourceGroups/RG/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/origins/origin2"))),
            com.azure.core.util.Context.NONE);
    }
}
