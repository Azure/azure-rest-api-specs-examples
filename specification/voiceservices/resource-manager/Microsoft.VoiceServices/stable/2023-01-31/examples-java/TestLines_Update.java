import com.azure.resourcemanager.voiceservices.models.TestLine;

/** Samples for TestLines Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/TestLines_Update.json
     */
    /**
     * Sample code: UpdateTestLineResource.
     *
     * @param manager Entry point to VoiceservicesManager.
     */
    public static void updateTestLineResource(com.azure.resourcemanager.voiceservices.VoiceservicesManager manager) {
        TestLine resource =
            manager
                .testLines()
                .getWithResponse("testrg", "myname", "myline", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
