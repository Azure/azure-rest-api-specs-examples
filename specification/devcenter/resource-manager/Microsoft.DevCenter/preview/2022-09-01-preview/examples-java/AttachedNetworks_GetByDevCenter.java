import com.azure.core.util.Context;

/** Samples for AttachedNetworks GetByDevCenter. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/AttachedNetworks_GetByDevCenter.json
     */
    /**
     * Sample code: AttachedNetworks_GetByDevCenter.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksGetByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.attachedNetworks().getByDevCenterWithResponse("rg1", "Contoso", "network-uswest3", Context.NONE);
    }
}
