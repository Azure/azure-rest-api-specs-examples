import com.azure.core.util.Context;

/** Samples for AttachedNetworks ListByProject. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/AttachedNetworks_ListByProject.json
     */
    /**
     * Sample code: AttachedNetworks_ListByProject.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksListByProject(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.attachedNetworks().listByProject("rg1", "{projectName}", null, Context.NONE);
    }
}
