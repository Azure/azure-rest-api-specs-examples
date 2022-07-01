import com.azure.core.util.Context;

/** Samples for ResourceProvider SapAvailabilityZoneDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPAvailabilityZoneDetails_eastus.json
     */
    /**
     * Sample code: SAPAvailabilityZoneDetails_eastus.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPAvailabilityZoneDetailsEastus(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapAvailabilityZoneDetailsWithResponse("centralus", null, Context.NONE);
    }
}
