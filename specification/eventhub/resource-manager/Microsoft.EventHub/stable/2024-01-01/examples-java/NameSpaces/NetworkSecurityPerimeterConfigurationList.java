
/**
 * Samples for NetworkSecurityPerimeterConfiguration List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/
     * NetworkSecurityPerimeterConfigurationList.json
     */
    /**
     * Sample code: NamspaceNetworkSecurityPerimeterConfigurationList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        namspaceNetworkSecurityPerimeterConfigurationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNetworkSecurityPerimeterConfigurations()
            .listWithResponse("SDK-EventHub-4794", "sdk-Namespace-5828", com.azure.core.util.Context.NONE);
    }
}
