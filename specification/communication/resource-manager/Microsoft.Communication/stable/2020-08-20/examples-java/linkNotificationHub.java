import com.azure.core.util.Context;
import com.azure.resourcemanager.communication.models.LinkNotificationHubParameters;

/** Samples for CommunicationService LinkNotificationHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/linkNotificationHub.json
     */
    /**
     * Sample code: Link notification hub.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void linkNotificationHub(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .linkNotificationHubWithResponse(
                "MyResourceGroup",
                "MyCommunicationResource",
                new LinkNotificationHubParameters()
                    .withResourceId(
                        "/subscriptions/12345/resourceGroups/MyOtherResourceGroup/providers/Microsoft.NotificationHubs/namespaces/MyNamespace/notificationHubs/MyHub")
                    .withConnectionString("Endpoint=sb://MyNamespace.servicebus.windows.net/;SharedAccessKey=abcd1234"),
                Context.NONE);
    }
}
