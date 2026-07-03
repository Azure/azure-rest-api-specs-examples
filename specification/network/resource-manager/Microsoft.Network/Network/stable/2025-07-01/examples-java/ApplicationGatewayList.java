
/**
 * Samples for ApplicationGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayList.json
     */
    /**
     * Sample code: Lists all application gateways in a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listsAllApplicationGatewaysInAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
