
/**
 * Samples for NetworkSecurityPerimeterConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/NetworkSecurityPerimeterConfigurationList.json
     */
    /**
     * Sample code: NamspaceNetworkSecurityPerimeterConfigurationList.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void namspaceNetworkSecurityPerimeterConfigurationList(
        com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().listWithResponse("SDK-EventHub-4794",
            "sdk-Namespace-5828", com.azure.core.util.Context.NONE);
    }
}
