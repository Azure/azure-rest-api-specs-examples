/** Samples for AttachedNetworks GetByProject. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/AttachedNetworks_GetByProject.json
     */
    /**
     * Sample code: AttachedNetworks_GetByProject.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void attachedNetworksGetByProject(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .attachedNetworks()
            .getByProjectWithResponse("rg1", "DevProject", "network-uswest3", com.azure.core.util.Context.NONE);
    }
}
