
/**
 * Samples for ManagedNetworkSettingsOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ManagedNetwork/listManagedNetworkV2.json
     */
    /**
     * Sample code: List ManagedNetworkSettings.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listManagedNetworkSettings(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedNetworkSettingsOperations().list("test-rg", "cognitive-account-name",
            com.azure.core.util.Context.NONE);
    }
}
