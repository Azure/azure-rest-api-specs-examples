/** Samples for VCenters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListVCenters.json
     */
    /**
     * Sample code: ListVCenters.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVCenters(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.vCenters().list(com.azure.core.util.Context.NONE);
    }
}
