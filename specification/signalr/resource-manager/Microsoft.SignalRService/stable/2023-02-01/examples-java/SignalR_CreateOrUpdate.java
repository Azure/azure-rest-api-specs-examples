
import com.azure.resourcemanager.signalr.models.AclAction;
import com.azure.resourcemanager.signalr.models.FeatureFlags;
import com.azure.resourcemanager.signalr.models.LiveTraceCategory;
import com.azure.resourcemanager.signalr.models.LiveTraceConfiguration;
import com.azure.resourcemanager.signalr.models.ManagedIdentity;
import com.azure.resourcemanager.signalr.models.ManagedIdentitySettings;
import com.azure.resourcemanager.signalr.models.ManagedIdentityType;
import com.azure.resourcemanager.signalr.models.NetworkAcl;
import com.azure.resourcemanager.signalr.models.PrivateEndpointAcl;
import com.azure.resourcemanager.signalr.models.ResourceSku;
import com.azure.resourcemanager.signalr.models.ServerlessSettings;
import com.azure.resourcemanager.signalr.models.ServerlessUpstreamSettings;
import com.azure.resourcemanager.signalr.models.ServiceKind;
import com.azure.resourcemanager.signalr.models.SignalRCorsSettings;
import com.azure.resourcemanager.signalr.models.SignalRFeature;
import com.azure.resourcemanager.signalr.models.SignalRNetworkACLs;
import com.azure.resourcemanager.signalr.models.SignalRRequestType;
import com.azure.resourcemanager.signalr.models.SignalRSkuTier;
import com.azure.resourcemanager.signalr.models.SignalRTlsSettings;
import com.azure.resourcemanager.signalr.models.UpstreamAuthSettings;
import com.azure.resourcemanager.signalr.models.UpstreamAuthType;
import com.azure.resourcemanager.signalr.models.UpstreamTemplate;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SignalR CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/signalr/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/SignalR_CreateOrUpdate
     * .json
     */
    /**
     * Sample code: SignalR_CreateOrUpdate.
     * 
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCreateOrUpdate(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().define("mySignalRService").withRegion("eastus").withExistingResourceGroup("myResourceGroup")
            .withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withSku(new ResourceSku().withName("Premium_P1").withTier(SignalRSkuTier.PREMIUM).withCapacity(1))
            .withKind(ServiceKind.SIGNALR)
            .withIdentity(new ManagedIdentity().withType(ManagedIdentityType.SYSTEM_ASSIGNED))
            .withTls(new SignalRTlsSettings().withClientCertEnabled(false))
            .withFeatures(Arrays.asList(
                new SignalRFeature().withFlag(FeatureFlags.SERVICE_MODE).withValue("Serverless")
                    .withProperties(mapOf()),
                new SignalRFeature().withFlag(FeatureFlags.ENABLE_CONNECTIVITY_LOGS).withValue("True")
                    .withProperties(mapOf()),
                new SignalRFeature().withFlag(FeatureFlags.ENABLE_MESSAGING_LOGS).withValue("False")
                    .withProperties(mapOf()),
                new SignalRFeature().withFlag(FeatureFlags.ENABLE_LIVE_TRACE).withValue("False")
                    .withProperties(mapOf())))
            .withLiveTraceConfiguration(new LiveTraceConfiguration().withEnabled("false").withCategories(
                Arrays.asList(new LiveTraceCategory().withName("ConnectivityLogs").withEnabled("true"))))
            .withCors(new SignalRCorsSettings().withAllowedOrigins(Arrays.asList("https://foo.com", "https://bar.com")))
            .withServerless(new ServerlessSettings().withConnectionTimeoutInSeconds(5))
            .withUpstream(new ServerlessUpstreamSettings().withTemplates(
                Arrays.asList(new UpstreamTemplate().withHubPattern("*").withEventPattern("connect,disconnect")
                    .withCategoryPattern("*").withUrlTemplate("https://example.com/chat/api/connect")
                    .withAuth(new UpstreamAuthSettings().withType(UpstreamAuthType.MANAGED_IDENTITY)
                        .withManagedIdentity(new ManagedIdentitySettings().withResource("api://example"))))))
            .withNetworkACLs(new SignalRNetworkACLs().withDefaultAction(AclAction.DENY)
                .withPublicNetwork(new NetworkAcl().withAllow(Arrays.asList(SignalRRequestType.CLIENT_CONNECTION)))
                .withPrivateEndpoints(Arrays
                    .asList(new PrivateEndpointAcl().withAllow(Arrays.asList(SignalRRequestType.SERVER_CONNECTION))
                        .withName("mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"))))
            .withPublicNetworkAccess("Enabled").withDisableLocalAuth(false).withDisableAadAuth(false).create();
    }

    // Use "Map.of" if available
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
