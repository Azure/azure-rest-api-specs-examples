/** Samples for CommunicationsGateways ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/CommunicationsGateways_ListByResourceGroup.json
     */
    /**
     * Sample code: ListCommunicationsGatewayResource.
     *
     * @param manager Entry point to VoiceservicesManager.
     */
    public static void listCommunicationsGatewayResource(
        com.azure.resourcemanager.voiceservices.VoiceservicesManager manager) {
        manager.communicationsGateways().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
