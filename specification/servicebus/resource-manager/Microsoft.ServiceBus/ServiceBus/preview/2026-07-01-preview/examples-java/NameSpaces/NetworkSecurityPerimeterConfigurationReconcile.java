
/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/NetworkSecurityPerimeterConfigurationReconcile.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurationList.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        networkSecurityPerimeterConfigurationList(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterConfigurations().reconcileWithResponse("SDK-ServiceBus-4794",
            "sdk-Namespace-5828", "resourceAssociation1", com.azure.core.util.Context.NONE);
    }
}
