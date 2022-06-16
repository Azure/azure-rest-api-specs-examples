import com.azure.resourcemanager.webpubsub.models.EventHandler;
import com.azure.resourcemanager.webpubsub.models.ManagedIdentitySettings;
import com.azure.resourcemanager.webpubsub.models.UpstreamAuthSettings;
import com.azure.resourcemanager.webpubsub.models.UpstreamAuthType;
import com.azure.resourcemanager.webpubsub.models.WebPubSubHubProperties;
import java.util.Arrays;

/** Samples for WebPubSubHubs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubHubs_CreateOrUpdate.json
     */
    /**
     * Sample code: WebPubSubHubs_CreateOrUpdate.
     *
     * @param manager Entry point to WebPubSubManager.
     */
    public static void webPubSubHubsCreateOrUpdate(com.azure.resourcemanager.webpubsub.WebPubSubManager manager) {
        manager
            .webPubSubHubs()
            .define("exampleHub")
            .withExistingWebPubSub("myResourceGroup", "myWebPubSubService")
            .withProperties(
                new WebPubSubHubProperties()
                    .withEventHandlers(
                        Arrays
                            .asList(
                                new EventHandler()
                                    .withUrlTemplate("http://host.com")
                                    .withUserEventPattern("*")
                                    .withSystemEvents(Arrays.asList("connect", "connected"))
                                    .withAuth(
                                        new UpstreamAuthSettings()
                                            .withType(UpstreamAuthType.MANAGED_IDENTITY)
                                            .withManagedIdentity(new ManagedIdentitySettings().withResource("abc"))))))
            .create();
    }
}
