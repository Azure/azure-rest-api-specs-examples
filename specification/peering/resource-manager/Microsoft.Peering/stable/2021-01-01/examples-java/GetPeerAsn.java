
/**
 * Samples for PeerAsns Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/GetPeerAsn.json
     */
    /**
     * Sample code: Get a peer ASN.
     * 
     * @param manager Entry point to PeeringManager.
     */
    public static void getAPeerASN(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.peerAsns().getWithResponse("peerAsnName", com.azure.core.util.Context.NONE);
    }
}
