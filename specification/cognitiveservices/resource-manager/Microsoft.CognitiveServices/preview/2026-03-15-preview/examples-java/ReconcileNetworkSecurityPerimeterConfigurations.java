
/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ReconcileNetworkSecurityPerimeterConfigurations.json
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
