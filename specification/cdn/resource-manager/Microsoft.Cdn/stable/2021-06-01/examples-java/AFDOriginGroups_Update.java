import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.AfdOriginGroupUpdateParameters;
import com.azure.resourcemanager.cdn.models.HealthProbeParameters;
import com.azure.resourcemanager.cdn.models.HealthProbeRequestType;
import com.azure.resourcemanager.cdn.models.LoadBalancingSettingsParameters;
import com.azure.resourcemanager.cdn.models.ProbeProtocol;

/** Samples for AfdOriginGroups Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Update.json
     */
    /**
     * Sample code: AFDOriginGroups_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginGroupsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdOriginGroups()
            .update(
                "RG",
                "profile1",
                "origingroup1",
                new AfdOriginGroupUpdateParameters()
                    .withLoadBalancingSettings(
                        new LoadBalancingSettingsParameters()
                            .withSampleSize(3)
                            .withSuccessfulSamplesRequired(3)
                            .withAdditionalLatencyInMilliseconds(1000))
                    .withHealthProbeSettings(
                        new HealthProbeParameters()
                            .withProbePath("/path2")
                            .withProbeRequestType(HealthProbeRequestType.NOT_SET)
                            .withProbeProtocol(ProbeProtocol.NOT_SET)
                            .withProbeIntervalInSeconds(10))
                    .withTrafficRestorationTimeToHealedOrNewEndpointsInMinutes(5),
                Context.NONE);
    }
}
