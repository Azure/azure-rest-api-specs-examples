/** Samples for AttachedNetworks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/AttachedNetworks_Create.json
     */
    /**
     * Sample code: AttachedNetworks_Create.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksCreate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .attachedNetworks()
            .define("{attachedNetworkConnectionName}")
            .withExistingDevcenter("rg1", "Contoso")
            .withNetworkConnectionId(
                "/subscriptions/{subscriptionId}/resourceGroups/rg1/providers/Microsoft.DevCenter/NetworkConnections/network-uswest3")
            .create();
    }
}
