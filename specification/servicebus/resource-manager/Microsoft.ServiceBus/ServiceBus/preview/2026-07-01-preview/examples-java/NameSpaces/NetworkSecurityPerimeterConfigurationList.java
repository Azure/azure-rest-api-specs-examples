
/**
 * Samples for NetworkSecurityPerimeterConfigurationOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/NetworkSecurityPerimeterConfigurationList.json
     */
    /**
     * Sample code: NamspaceNetworkSecurityPerimeterConfigurationList.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void namspaceNetworkSecurityPerimeterConfigurationList(
        com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurationOperations().list("SDK-ServiceBus-4794",
            "sdk-Namespace-5828", com.azure.core.util.Context.NONE);
    }
}
