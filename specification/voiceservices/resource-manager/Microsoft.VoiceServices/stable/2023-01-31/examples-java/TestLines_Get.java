/** Samples for TestLines Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/TestLines_Get.json
     */
    /**
     * Sample code: GetTestLineResource.
     *
     * @param manager Entry point to VoiceServicesManager.
     */
    public static void getTestLineResource(com.azure.resourcemanager.voiceservices.VoiceServicesManager manager) {
        manager.testLines().getWithResponse("testrg", "myname", "myline", com.azure.core.util.Context.NONE);
    }
}
