
/**
 * Samples for NetworkSecurityPerimeterConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/NetworkSecurityPerimeterConfigurationReconcile.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationList.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void
        networkSecurityPerimeterConfigurationList(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().createOrUpdate("SDK-EventHub-4794",
            "sdk-Namespace-5828", "resourceAssociation1", com.azure.core.util.Context.NONE);
    }
}
