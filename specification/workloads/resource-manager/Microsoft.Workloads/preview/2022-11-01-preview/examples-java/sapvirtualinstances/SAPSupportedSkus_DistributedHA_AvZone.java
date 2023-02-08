/** Samples for ResourceProvider SapSupportedSku. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPSupportedSkus_DistributedHA_AvZone.json
     */
    /**
     * Sample code: SAPSupportedSkus_DistributedHA_AvZone.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSupportedSkusDistributedHAAvZone(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapSupportedSkuWithResponse("centralus", null, com.azure.core.util.Context.NONE);
    }
}
