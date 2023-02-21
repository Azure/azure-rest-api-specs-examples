import com.azure.resourcemanager.voiceservices.models.TestLinePurpose;

/** Samples for TestLines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/TestLines_CreateOrUpdate.json
     */
    /**
     * Sample code: CreateTestLineResource.
     *
     * @param manager Entry point to VoiceservicesManager.
     */
    public static void createTestLineResource(com.azure.resourcemanager.voiceservices.VoiceservicesManager manager) {
        manager
            .testLines()
            .define("myline")
            .withRegion("useast")
            .withExistingCommunicationsGateway("testrg", "myname")
            .withPhoneNumber("+1-555-1234")
            .withPurpose(TestLinePurpose.AUTOMATED)
            .create();
    }
}
