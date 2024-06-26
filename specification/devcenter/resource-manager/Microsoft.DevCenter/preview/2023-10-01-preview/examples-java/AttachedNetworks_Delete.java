/** Samples for AttachedNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/AttachedNetworks_Delete.json
     */
    /**
     * Sample code: AttachedNetworks_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.attachedNetworks().delete("rg1", "Contoso", "network-uswest3", com.azure.core.util.Context.NONE);
    }
}
