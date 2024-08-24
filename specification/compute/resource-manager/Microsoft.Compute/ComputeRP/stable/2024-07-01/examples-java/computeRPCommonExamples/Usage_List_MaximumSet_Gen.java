
/**
 * Samples for Usage List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * computeRPCommonExamples/Usage_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Usage_List_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void usageListMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getUsages().list("4_.", com.azure.core.util.Context.NONE);
    }
}
