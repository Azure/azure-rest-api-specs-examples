import com.azure.core.util.Context;

/** Samples for AvailabilitySets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListAvailabilitySetsBySubscription.json
     */
    /**
     * Sample code: ListAvailabilitySetsBySubscription.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listAvailabilitySetsBySubscription(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.availabilitySets().list(Context.NONE);
    }
}
