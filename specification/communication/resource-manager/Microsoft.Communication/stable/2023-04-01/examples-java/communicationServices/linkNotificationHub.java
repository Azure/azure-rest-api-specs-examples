
import com.azure.resourcemanager.communication.models.LinkNotificationHubParameters;

/**
 * Samples for CommunicationServices LinkNotificationHub.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2023-04-01/examples/
     * communicationServices/linkNotificationHub.json
     */
    /**
     * Sample code: Link notification hub.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void linkNotificationHub(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().linkNotificationHubWithResponse("MyResourceGroup", "MyCommunicationResource",
            new LinkNotificationHubParameters().withResourceId(
                "/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyOtherResourceGroup/providers/Microsoft.NotificationHubs/namespaces/MyNamespace/notificationHubs/MyHub")
                .withConnectionString("Endpoint=sb://MyNamespace.servicebus.windows.net/;SharedAccessKey=abcd1234"),
            com.azure.core.util.Context.NONE);
    }
}
