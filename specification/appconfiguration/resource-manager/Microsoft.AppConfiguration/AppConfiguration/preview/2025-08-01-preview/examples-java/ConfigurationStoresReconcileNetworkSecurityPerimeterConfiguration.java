
/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/ConfigurationStoresReconcileNetworkSecurityPerimeterConfiguration.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurations_Reconcile.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void networkSecurityPerimeterConfigurationsReconcile(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.networkSecurityPerimeterConfigurations().reconcile("myResourceGroup", "contoso",
            "804a12bb-1349-4228-81be-8fe888aae04e.myAssociationName", com.azure.core.util.Context.NONE);
    }
}
