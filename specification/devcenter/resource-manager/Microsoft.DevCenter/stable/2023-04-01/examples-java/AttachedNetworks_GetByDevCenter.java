/** Samples for AttachedNetworks GetByDevCenter. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/AttachedNetworks_GetByDevCenter.json
     */
    /**
     * Sample code: AttachedNetworks_GetByDevCenter.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksGetByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .attachedNetworks()
            .getByDevCenterWithResponse("rg1", "Contoso", "network-uswest3", com.azure.core.util.Context.NONE);
    }
}
