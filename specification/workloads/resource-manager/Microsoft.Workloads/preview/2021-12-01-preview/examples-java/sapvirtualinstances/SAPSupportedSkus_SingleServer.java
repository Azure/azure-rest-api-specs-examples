import com.azure.core.util.Context;

/** Samples for ResourceProvider SapSupportedSku. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPSupportedSkus_SingleServer.json
     */
    /**
     * Sample code: SAPSupportedSkus_SingleServer.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPSupportedSkusSingleServer(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.resourceProviders().sapSupportedSkuWithResponse("centralus", null, Context.NONE);
    }
}
