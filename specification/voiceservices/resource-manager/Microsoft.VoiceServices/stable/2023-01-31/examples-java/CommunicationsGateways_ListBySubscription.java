/** Samples for CommunicationsGateways List. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/CommunicationsGateways_ListBySubscription.json
     */
    /**
     * Sample code: ListCommunicationsGatewayResourceSub.
     *
     * @param manager Entry point to VoiceServicesManager.
     */
    public static void listCommunicationsGatewayResourceSub(
        com.azure.resourcemanager.voiceservices.VoiceServicesManager manager) {
        manager.communicationsGateways().list(com.azure.core.util.Context.NONE);
    }
}
