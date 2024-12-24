
/**
 * Samples for MobileNetworks ListSimGroups.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * SimGroupListByMobileNetwork.json
     */
    /**
     * Sample code: List SIM groups in a mobile network.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void
        listSIMGroupsInAMobileNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.mobileNetworks().listSimGroups("rg1", "testMobileNetwork", com.azure.core.util.Context.NONE);
    }
}
