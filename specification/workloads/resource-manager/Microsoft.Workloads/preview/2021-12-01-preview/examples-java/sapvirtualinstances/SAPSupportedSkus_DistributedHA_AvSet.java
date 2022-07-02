import com.azure.core.util.Context;

/** Samples for ResourceProvider SapSupportedSku. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPSupportedSkus_DistributedHA_AvSet.json
     */
    /**
     * Sample code: SAPSupportedSkus_DistributedHA_AvSet.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSupportedSkusDistributedHAAvSet(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapSupportedSkuWithResponse("centralus", null, Context.NONE);
    }
}
