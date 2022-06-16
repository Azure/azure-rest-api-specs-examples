import com.azure.core.util.Context;

/** Samples for VmmServers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVMMServersBySubscription.json
     */
    /**
     * Sample code: ListVmmServersBySubscription.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVmmServersBySubscription(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().list(Context.NONE);
    }
}
