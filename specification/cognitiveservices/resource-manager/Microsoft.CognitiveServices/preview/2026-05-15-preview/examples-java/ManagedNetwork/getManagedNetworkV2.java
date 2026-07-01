
/**
 * Samples for ManagedNetworkSettingsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ManagedNetwork/getManagedNetworkV2.json
     */
    /**
     * Sample code: Get ManagedNetworkSettings.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getManagedNetworkSettings(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedNetworkSettingsOperations().getWithResponse("test-rg", "cognitive-account-name", "default",
            com.azure.core.util.Context.NONE);
    }
}
