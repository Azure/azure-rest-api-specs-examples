import com.azure.core.util.Context;

/** Samples for Usage List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/Usage_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Usage_List_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void usageListMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getUsages().list("_--", Context.NONE);
    }
}
