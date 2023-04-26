import com.azure.resourcemanager.voiceservices.models.CommunicationsGateway;

/** Samples for CommunicationsGateways Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/CommunicationsGateways_Update.json
     */
    /**
     * Sample code: UpdateCommunicationsGatewayResource.
     *
     * @param manager Entry point to VoiceServicesManager.
     */
    public static void updateCommunicationsGatewayResource(
        com.azure.resourcemanager.voiceservices.VoiceServicesManager manager) {
        CommunicationsGateway resource =
            manager
                .communicationsGateways()
                .getByResourceGroupWithResponse("testrg", "myname", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
