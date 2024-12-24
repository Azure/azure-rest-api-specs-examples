
/**
 * Samples for RoutingInfo Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * RoutingInfoPacketCoreControlPlane.json
     */
    /**
     * Sample code: Get routing information for the packet core.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void
        getRoutingInformationForThePacketCore(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.routingInfoes().getWithResponse("rg1", "TestPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
