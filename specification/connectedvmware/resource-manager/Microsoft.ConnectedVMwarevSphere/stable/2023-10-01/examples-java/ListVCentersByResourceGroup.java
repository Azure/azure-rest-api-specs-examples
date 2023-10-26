/** Samples for VCenters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListVCentersByResourceGroup.json
     */
    /**
     * Sample code: ListVCentersByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVCentersByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.vCenters().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
