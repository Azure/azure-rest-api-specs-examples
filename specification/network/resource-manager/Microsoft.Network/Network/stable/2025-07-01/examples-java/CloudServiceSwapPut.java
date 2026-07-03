
import com.azure.resourcemanager.network.fluent.models.SwapResourceInner;
import com.azure.resourcemanager.network.models.SlotType;
import com.azure.resourcemanager.network.models.SwapResourceProperties;

/**
 * Samples for VipSwap Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServiceSwapPut.json
     */
    /**
     * Sample code: Put vip swap operation.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void putVipSwapOperation(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVipSwaps().create("rg1", "testCloudService",
            new SwapResourceInner().withProperties(new SwapResourceProperties().withSlotType(SlotType.PRODUCTION)),
            com.azure.core.util.Context.NONE);
    }
}
