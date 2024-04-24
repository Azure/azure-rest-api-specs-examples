
/**
 * Samples for AttachedNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/AttachedNetworks_Create.
     * json
     */
    /**
     * Sample code: AttachedNetworks_Create.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksCreate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.attachedNetworks().define("network-uswest3").withExistingDevcenter("rg1", "Contoso")
            .withNetworkConnectionId(
                "/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/NetworkConnections/network-uswest3")
            .create();
    }
}
