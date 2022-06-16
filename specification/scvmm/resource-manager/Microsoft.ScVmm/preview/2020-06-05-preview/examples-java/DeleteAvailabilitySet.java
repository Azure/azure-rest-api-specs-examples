import com.azure.core.util.Context;

/** Samples for AvailabilitySets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteAvailabilitySet.json
     */
    /**
     * Sample code: DeleteAvailabilitySet.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void deleteAvailabilitySet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.availabilitySets().delete("testrg", "HRAvailabilitySet", null, Context.NONE);
    }
}
