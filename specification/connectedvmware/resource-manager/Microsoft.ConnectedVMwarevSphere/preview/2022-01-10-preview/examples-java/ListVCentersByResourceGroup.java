import com.azure.core.util.Context;

/** Samples for VCenters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVCentersByResourceGroup.json
     */
    /**
     * Sample code: ListVCentersByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVCentersByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.vCenters().listByResourceGroup("testrg", Context.NONE);
    }
}
