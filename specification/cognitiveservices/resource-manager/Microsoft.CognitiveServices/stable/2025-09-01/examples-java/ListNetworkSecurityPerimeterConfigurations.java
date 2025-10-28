
/**
 * Samples for NetworkSecurityPerimeterConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ListNetworkSecurityPerimeterConfigurations.json
     */
    /**
     * Sample code: ListNetworkSecurityPerimeterConfigurations.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listNetworkSecurityPerimeterConfigurations(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.networkSecurityPerimeterConfigurations().list("resourceGroupName", "accountName",
            com.azure.core.util.Context.NONE);
    }
}
