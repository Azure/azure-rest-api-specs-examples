
/**
 * Samples for RoutingInfo List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * RoutingInfoListPacketCoreControlPlane.json
     */
    /**
     * Sample code: List routing information for the packet core.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void
        listRoutingInformationForThePacketCore(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.routingInfoes().list("rg1", "TestPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
