import com.azure.core.util.Context;

/** Samples for VCenters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteVCenter.json
     */
    /**
     * Sample code: DeleteVCenter.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteVCenter(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.vCenters().delete("testrg", "ContosoVCenter", null, Context.NONE);
    }
}
