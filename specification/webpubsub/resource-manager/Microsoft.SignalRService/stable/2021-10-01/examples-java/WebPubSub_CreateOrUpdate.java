import com.azure.resourcemanager.webpubsub.models.AclAction;
import com.azure.resourcemanager.webpubsub.models.LiveTraceCategory;
import com.azure.resourcemanager.webpubsub.models.LiveTraceConfiguration;
import com.azure.resourcemanager.webpubsub.models.ManagedIdentity;
import com.azure.resourcemanager.webpubsub.models.ManagedIdentityType;
import com.azure.resourcemanager.webpubsub.models.NetworkAcl;
import com.azure.resourcemanager.webpubsub.models.PrivateEndpointAcl;
import com.azure.resourcemanager.webpubsub.models.ResourceSku;
import com.azure.resourcemanager.webpubsub.models.WebPubSubNetworkACLs;
import com.azure.resourcemanager.webpubsub.models.WebPubSubRequestType;
import com.azure.resourcemanager.webpubsub.models.WebPubSubSkuTier;
import com.azure.resourcemanager.webpubsub.models.WebPubSubTlsSettings;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for WebPubSub CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_CreateOrUpdate.json
     */
    /**
     * Sample code: WebPubSub_CreateOrUpdate.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubCreateOrUpdate(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubs()
            .define("myWebPubSubService")
            .withRegion("eastus")
            .withExistingResourceGroup("myResourceGroup")
            .withTags(mapOf("key1", "value1"))
            .withSku(new ResourceSku().withName("Standard_S1").withTier(WebPubSubSkuTier.STANDARD).withCapacity(1))
            .withIdentity(new ManagedIdentity().withType(ManagedIdentityType.SYSTEM_ASSIGNED))
            .withTls(new WebPubSubTlsSettings().withClientCertEnabled(false))
            .withLiveTraceConfiguration(
                new LiveTraceConfiguration()
                    .withEnabled("false")
                    .withCategories(
                        Arrays.asList(new LiveTraceCategory().withName("ConnectivityLogs").withEnabled("true"))))
            .withNetworkACLs(
                new WebPubSubNetworkACLs()
                    .withDefaultAction(AclAction.DENY)
                    .withPublicNetwork(
                        new NetworkAcl().withAllow(Arrays.asList(WebPubSubRequestType.CLIENT_CONNECTION)))
                    .withPrivateEndpoints(
                        Arrays
                            .asList(
                                new PrivateEndpointAcl()
                                    .withAllow(Arrays.asList(WebPubSubRequestType.SERVER_CONNECTION))
                                    .withName("mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e"))))
            .withPublicNetworkAccess("Enabled")
            .withDisableLocalAuth(false)
            .withDisableAadAuth(false)
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
