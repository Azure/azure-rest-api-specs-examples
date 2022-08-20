import com.azure.core.util.Context;

/** Samples for VCenters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVCenters.json
     */
    /**
     * Sample code: ListVCenters.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVCenters(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.vCenters().list(Context.NONE);
    }
}
