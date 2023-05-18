/** Samples for AttachedNetworks ListByDevCenter. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/AttachedNetworks_ListByDevCenter.json
     */
    /**
     * Sample code: AttachedNetworks_ListByDevCenter.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksListByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.attachedNetworks().listByDevCenter("rg1", "Contoso", null, com.azure.core.util.Context.NONE);
    }
}
