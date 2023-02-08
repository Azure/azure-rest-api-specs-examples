/** Samples for ResourceProvider SapAvailabilityZoneDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPAvailabilityZoneDetails_northeurope.json
     */
    /**
     * Sample code: SAPAvailabilityZoneDetails_northeurope.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPAvailabilityZoneDetailsNortheurope(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .resourceProviders()
            .sapAvailabilityZoneDetailsWithResponse("centralus", null, com.azure.core.util.Context.NONE);
    }
}
