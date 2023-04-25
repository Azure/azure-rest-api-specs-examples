/** Samples for CommunicationsGateways Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/CommunicationsGateways_Delete.json
     */
    /**
     * Sample code: DeleteCommunicationsGatewayResource.
     *
     * @param manager Entry point to VoiceServicesManager.
     */
    public static void deleteCommunicationsGatewayResource(
        com.azure.resourcemanager.voiceservices.VoiceServicesManager manager) {
        manager.communicationsGateways().delete("testrg", "myname", com.azure.core.util.Context.NONE);
    }
}
