import com.azure.core.util.Context;

/** Samples for VCenters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetVCenter.json
     */
    /**
     * Sample code: GetVCenter.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getVCenter(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.vCenters().getByResourceGroupWithResponse("testrg", "ContosoVCenter", Context.NONE);
    }
}
