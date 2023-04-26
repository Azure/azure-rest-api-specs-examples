/** Samples for TestLines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/TestLines_Delete.json
     */
    /**
     * Sample code: DeleteTestLineResource.
     *
     * @param manager Entry point to VoiceServicesManager.
     */
    public static void deleteTestLineResource(com.azure.resourcemanager.voiceservices.VoiceServicesManager manager) {
        manager.testLines().delete("testrg", "myname", "myline", com.azure.core.util.Context.NONE);
    }
}
