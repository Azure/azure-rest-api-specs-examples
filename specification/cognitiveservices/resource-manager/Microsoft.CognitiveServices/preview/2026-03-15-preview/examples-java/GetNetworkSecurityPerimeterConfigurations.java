
/**
 * Samples for NetworkSecurityPerimeterConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/GetNetworkSecurityPerimeterConfigurations.json
     */
    /**
     * Sample code: GetNetworkSecurityPerimeterConfigurations.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getNetworkSecurityPerimeterConfigurations(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.networkSecurityPerimeterConfigurations().getWithResponse("resourceGroupName", "accountName",
            "NSPConfigurationName", com.azure.core.util.Context.NONE);
    }
}
