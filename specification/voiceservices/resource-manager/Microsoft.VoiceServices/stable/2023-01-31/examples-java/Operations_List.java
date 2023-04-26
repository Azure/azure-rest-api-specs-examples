/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/voiceservices/resource-manager/Microsoft.VoiceServices/stable/2023-01-31/examples/Operations_List.json
     */
    /**
     * Sample code: OperationsList.
     *
     * @param manager Entry point to VoiceServicesManager.
     */
    public static void operationsList(com.azure.resourcemanager.voiceservices.VoiceServicesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
