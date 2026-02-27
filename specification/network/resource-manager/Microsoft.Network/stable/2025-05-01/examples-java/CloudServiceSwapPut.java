
import com.azure.resourcemanager.network.fluent.models.SwapResourceInner;
import com.azure.resourcemanager.network.models.SlotType;
import com.azure.resourcemanager.network.models.SwapResourceProperties;

/**
 * Samples for VipSwap Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/CloudServiceSwapPut.json
     */
    /**
     * Sample code: Put vip swap operation.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putVipSwapOperation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVipSwaps().create("rg1", "testCloudService",
            new SwapResourceInner().withProperties(new SwapResourceProperties().withSlotType(SlotType.PRODUCTION)),
            com.azure.core.util.Context.NONE);
    }
}
