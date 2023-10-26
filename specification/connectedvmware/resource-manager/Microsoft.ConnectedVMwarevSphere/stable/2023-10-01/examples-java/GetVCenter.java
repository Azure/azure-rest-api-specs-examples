/** Samples for VCenters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/GetVCenter.json
     */
    /**
     * Sample code: GetVCenter.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getVCenter(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.vCenters().getByResourceGroupWithResponse("testrg", "ContosoVCenter", com.azure.core.util.Context.NONE);
    }
}
