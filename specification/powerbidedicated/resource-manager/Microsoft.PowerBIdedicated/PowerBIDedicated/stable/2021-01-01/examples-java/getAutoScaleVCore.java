
/**
 * Samples for AutoScaleVCores GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-01-01/getAutoScaleVCore.json
     */
    /**
     * Sample code: Get details of an auto scale v-core.
     * 
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void
        getDetailsOfAnAutoScaleVCore(com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.autoScaleVCores().getByResourceGroupWithResponse("TestRG", "testvcore",
            com.azure.core.util.Context.NONE);
    }
}
