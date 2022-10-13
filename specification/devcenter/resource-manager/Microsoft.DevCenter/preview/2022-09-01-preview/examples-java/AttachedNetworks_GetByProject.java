import com.azure.core.util.Context;

/** Samples for AttachedNetworks GetByProject. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/AttachedNetworks_GetByProject.json
     */
    /**
     * Sample code: AttachedNetworks_GetByProject.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksGetByProject(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.attachedNetworks().getByProjectWithResponse("rg1", "{projectName}", "network-uswest3", Context.NONE);
    }
}
