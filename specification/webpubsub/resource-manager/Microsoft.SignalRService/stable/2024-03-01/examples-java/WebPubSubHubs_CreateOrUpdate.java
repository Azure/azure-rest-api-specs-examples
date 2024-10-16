
import com.azure.resourcemanager.webpubsub.models.EventHandler;
import com.azure.resourcemanager.webpubsub.models.EventHubEndpoint;
import com.azure.resourcemanager.webpubsub.models.EventListener;
import com.azure.resourcemanager.webpubsub.models.EventNameFilter;
import com.azure.resourcemanager.webpubsub.models.ManagedIdentitySettings;
import com.azure.resourcemanager.webpubsub.models.UpstreamAuthSettings;
import com.azure.resourcemanager.webpubsub.models.UpstreamAuthType;
import com.azure.resourcemanager.webpubsub.models.WebPubSubHubProperties;
import java.util.Arrays;

/**
 * Samples for WebPubSubHubs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2024-03-01/examples/
     * WebPubSubHubs_CreateOrUpdate.json
     */
    /**
     * Sample code: WebPubSubHubs_CreateOrUpdate.
     * 
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubHubsCreateOrUpdate(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubHubs().define(
                "exampleHub")
            .withExistingWebPubSub("myResourceGroup",
                "myWebPubSubService")
            .withProperties(new WebPubSubHubProperties()
                .withEventHandlers(Arrays.asList(new EventHandler().withUrlTemplate("http://host.com")
                    .withUserEventPattern("*").withSystemEvents(Arrays.asList("connect", "connected"))
                    .withAuth(new UpstreamAuthSettings().withType(UpstreamAuthType.MANAGED_IDENTITY)
                        .withManagedIdentity(new ManagedIdentitySettings().withResource("abc")))))
                .withEventListeners(
                    Arrays
                        .asList(new EventListener()
                            .withFilter(new EventNameFilter()
                                .withSystemEvents(Arrays.asList("connected", "disconnected")).withUserEventPattern("*"))
                            .withEndpoint(
                                new EventHubEndpoint().withFullyQualifiedNamespace("example.servicebus.windows.net")
                                    .withEventHubName("eventHubName1"))))
                .withAnonymousConnectPolicy("allow").withWebSocketKeepAliveIntervalInSeconds(50))
            .create();
    }
}
