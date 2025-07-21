
/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ReconcileNetworkSecurityPerimeterConfigurations.json
     */
    /**
     * Sample code: ReconcileNetworkSecurityPerimeterConfigurations.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void reconcileNetworkSecurityPerimeterConfigurations(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.networkSecurityPerimeterConfigurations().reconcile("resourceGroupName", "accountName",
            "NSPConfigurationName", com.azure.core.util.Context.NONE);
    }
}
