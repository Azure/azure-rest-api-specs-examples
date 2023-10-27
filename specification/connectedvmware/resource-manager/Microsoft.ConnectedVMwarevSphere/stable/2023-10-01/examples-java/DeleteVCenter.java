/** Samples for VCenters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteVCenter.json
     */
    /**
     * Sample code: DeleteVCenter.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteVCenter(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.vCenters().delete("testrg", "ContosoVCenter", null, com.azure.core.util.Context.NONE);
    }
}
