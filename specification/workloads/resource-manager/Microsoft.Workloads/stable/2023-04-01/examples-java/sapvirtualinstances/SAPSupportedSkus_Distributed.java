/** Samples for ResourceProvider SapSupportedSku. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPSupportedSkus_Distributed.json
     */
    /**
     * Sample code: SAPSupportedSkus_Distributed.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSupportedSkusDistributed(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapSupportedSkuWithResponse("centralus", null, com.azure.core.util.Context.NONE);
    }
}
