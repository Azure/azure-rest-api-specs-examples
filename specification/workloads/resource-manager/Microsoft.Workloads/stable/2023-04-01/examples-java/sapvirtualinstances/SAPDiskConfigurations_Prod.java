/** Samples for ResourceProvider SapDiskConfigurations. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPDiskConfigurations_Prod.json
     */
    /**
     * Sample code: SAPDiskConfigurations_Prod.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPDiskConfigurationsProd(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .resourceProviders()
            .sapDiskConfigurationsWithResponse("centralus", null, com.azure.core.util.Context.NONE);
    }
}
