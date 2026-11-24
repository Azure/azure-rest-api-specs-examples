
/**
 * Samples for NetworkSecurityPerimeterConfigurations GetResourceAssociationName.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/NetworkSecurityPerimeterConfigurationAssociationproxy.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationassociationProxyName.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void networkSecurityPerimeterConfigurationassociationProxyName(
        com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().getResourceAssociationNameWithResponse(
            "SDK-ServiceBus-4794", "sdk-Namespace-5828", "resourceAssociation1", com.azure.core.util.Context.NONE);
    }
}
