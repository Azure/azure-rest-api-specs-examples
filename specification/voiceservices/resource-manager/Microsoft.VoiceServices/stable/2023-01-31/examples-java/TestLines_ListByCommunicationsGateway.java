/** Samples for TestLines ListByCommunicationsGateway. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/TestLines_ListByCommunicationsGateway.json
     */
    /**
     * Sample code: ListTestLineResource.
     *
     * @param manager Entry point to VoiceServicesManager.
     */
    public static void listTestLineResource(com.azure.resourcemanager.voiceservices.VoiceServicesManager manager) {
        manager.testLines().listByCommunicationsGateway("testrg", "myname", com.azure.core.util.Context.NONE);
    }
}
